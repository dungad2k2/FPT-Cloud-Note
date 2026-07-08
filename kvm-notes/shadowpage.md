# Some note about shadow pages using by KVM 
This articles originate from this CVE https://vexxhost.com/blog/cve-2026-53359-openstack-kvm-x86-compute-isolation/ . This CVE is about Linux KVM/x86 vulnerability in KVM's x86 MNU implementation. By exploit that CVE, untrusted guest can escape tho the host in a KVM/x86 environment 
## Normal memory translation inside a VM 
A process inside the guest uses a **guest virtual address**: 
```
Guest process VA
   -> guest page table

Guest physical address
   -> KVM / hardware second-level translation

Host physical address 
```

In this article, we focus on KVM second level translation from Guest physical address to Host physical address 

## Shadow paging what is it ?

Before EPT/NPT (Technology that help hardware can walk both guest page tables and host second-level page tables ), KVM cannot let the guest directly control the page tables used by the real CPU. So KVM create **show page tables**. 

KVM builds a host-controlled page table that represents the combined translation from **Guest Virtual Address** to **Host Physical Address**.

So that CPU uses a page table created by KVM, not the guest's real page table. It called **show page table**. 

Each 4 KB page that stores entries in this shadow page table is a shadow page. In KVM code, these are represented by structures such as struct `kvm_mmu_page`, and their entries are often called SPTEs, or Shadow Page Table Entries. The Linux KVM documentation describes the x86 shadow MMU as maintaining these shadow pages and their roles/metadata.

MMU emulation make the guest OS thinks it owns the MMU. That means KVM watches or traps guest MMU operations and updates its own shadow page tables accordingly.

```
Guest modifies its PTE:
    "map virtual address X to guest physical page Y"

KVM must update shadow page table:
    "map virtual address X to host physical page Z"
```
## Shadow MMU and how does it still alive beside hardware TDP MMU 

With hardware EPT/NPT:

```
Guest page table: GVA -> GPA
EPT/NPT:          GPA -> HPA
```

With shadow paging:

```
Shadow page table: GVA -> HPA
```

Shadow MMU logic is still relevant even on modern CPU, especially with **neted virtualization**. In nested virtualization, KVM may need to synthesize or shadow complex translations involving L2 guest page tables and L1-controlled EPT/NPT tables. In this CVE above, KVM shadow paging issue notes that exploitation is relevant when x86 nested virtualization is enabled when shadow paging is used directly with EPT/NPT disabled. 

## Why shadow MMU make host kernel corrupted

KVM shadow pages live in host kernel memory. If KVM frees a shadows page but still has a stale pointer or stale reverse map entry to it, later MMU activity may access memory that has already been freed that is a classic **use-after-free**: 
```
1. KVM allocates shadow page
2. Guest activity causes KVM to free/zap it
3. KVM still has stale reference to old shadow page
4. Later guest-triggered MMU path uses that stale reference
5. Host kernel memory can be corrupted
```

The guest controls its own page tables. It can legally do many MMU operations:

```
map page
unmap page
change permissions
switch CR3
invalidate TLB
cause page faults
run nested guest page-table activity
```

But if KVM’s shadow MMU bookkeeping has a bug, then guest-controlled page-table changes can drive KVM into bad internal states. This make host kernel memory corruption