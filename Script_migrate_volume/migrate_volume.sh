#!/bin/bash

# Read the volume id list from file
volumeid_list=$(cat volumeid_list.txt)
volumeid_array=($volumeid_list)
total_volumes=${#volumeid_array[@]}
migrating_volumes=()
concurrent_migrations=1

destination_volumes=("sgn09san20_pro_ops_vol01" "sgn09san20_pro_ops_vol02")
destination_index=0

# Function to check the status of migrating volumes
check_status() {
    for i in "${!migrating_volumes[@]}"; do
        checkVolumeid=${migrating_volumes[$i]}
        migration_status=$(openstack volume show $checkVolumeid -f value -c migration_status)
        if [ "$migration_status" != "migrating" ] && [ "$migration_status" != "starting" ]; then
            #sed -i "s/$checkVolumeid/$checkVolumeid $migration_status/" volumeid_list.txt
            echo "Volume $checkVolumeid migration status: $migration_status"
            count_snap_shot=$(openstack volume snapshot list --all-projects --volume $checkVolumeid -f value -c ID | wc -l)
            if [ "$count_snap_shot" -gt 0 ]; then
                    sed -i "s/$checkVolumeid/$checkVolumeid $migration_status $count_snap_shot/" volumeid_list.txt
            else
                    sed -i "s/$checkVolumeid/$checkVolumeid $migration_status/" volumeid_list.txt
            fi
            if [ "${original_status[$checkVolumeid]}" == "in-use" ]; then
                openstack volume set --state in-use $checkVolumeid
                echo "Volume $checkVolumeid state set back to in-use"
            fi
            unset migrating_volumes[$i]
        fi
    done
    migrating_volumes=("${migrating_volumes[@]}") # Re-index array
}

declare -A original_status
declare -A vm_id
# Iterate through the volume id list and migrate $concurrent_migrations volumes concurrently
for volumeid in "${volumeid_array[@]}"; do
    volume_status=$(openstack volume show $volumeid -f value -c status)
    current_host=$(openstack volume show $volumeid -f value -c os-vol-host-attr:host | cut -d'@' -f1)
  #  serverid=$(openstack volume show $volumeid -f json -c attachments | jq -r '.attachments[0].server_id')

    if [ "$volume_status" == "in-use" ]; then
        openstack volume set --state available $volumeid
        original_status[$volumeid]="in-use"
        echo "Volume $volumeid state set to available"
    elif [ "$volume_status" != "available" ]; then
        echo "Skipping volume $volumeid with status $volume_status"
        continue
    fi
    if [ -n "$serverid" ] && [ "$serverid" != "null" ]; then
         vm_state=$(openstack server show $serverid -f value -c status)
         vm_id[$volumeid]=$serverid
         if [ "$vm_state" == "SHUTOFF" ]; then
             sed -i "s/$checkVolumeid/$checkVolumeid Server SHUTOFF/" volumeid_list.txt
             if [ $original_status[$volumeid] == "in-use" ]; then
                openstack volume set --state in-use $volumeid
                echo "Volume $volumeid state set back to in-use"
             fi
             echo "VM $serverid SHUTOFF. Skipping..."
             continue
         fi
    fi

    while [ ${#migrating_volumes[@]} -ge $concurrent_migrations ]; do
        check_status
        sleep 5
    done

    cinder migrate --force-host-copy True $volumeid ${current_host}@PlatinumSSD_4#${destination_volumes[$destination_index]} &
    echo "Migrating volume $volumeid to ${current_host}@PlatinumSSD_4#${destination_volumes[$destination_index]}"
    migrating_volumes+=($volumeid)
    destination_index=$(( (destination_index + 1) % 2 ))
done

# Final check to ensure all migrations are complete
while [ ${#migrating_volumes[@]} -gt 0 ]; do
    check_status
    sleep 5
done

echo "All volumes migrated"
