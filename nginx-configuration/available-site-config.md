# Available-site-config 
### Step 1: Create an directory in `/var/www/yourwebsite`
### Step 2: Import html, css file in this directory that created
### Step 3: Config file in path  `/etc/nginx/sites-available/` 
```
##yourwebsite:namefile
server {
  listen 8000 ssl;


  server_name example.com;

  ssl_certificate     /home/buidung/cert/example.com.crt;
  ssl_certificate_key /home/buidung/cert/example.com.key;
  ssl_protocols       TLSv1 TLSv1.1 TLSv1.2 ;
  ssl_ciphers         HIGH:!aNULL:!MD5;

  location / {
    root /var/www/web3;
    index index.html index.htm;
  }
}

```
### Step 4: Enabled virtual host in `sites-enabled` by this command
```
sudo ln -s /etc/nginx/sites-available/yourwebsite /etc/nginx/sites-enabled
```
### Step 5: Restart service nginx