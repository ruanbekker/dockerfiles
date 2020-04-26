```
apt install geoip-bin geoip-database-extra libgeoip1 libnginx-mod-http-geoip geoip-database-extra libgeoip1 libnginx-mod-http-geoip nginx-module-geoip -y
find / -name ngx_http_geoip_module.so
mkdir /etc/nginx/geoip
cd /etc/nginx/geoip
wget https://mirrors-cdn.liferay.com/geolite.maxmind.com/download/geoip/database/GeoIP.dat.gz
wget https://mirrors-cdn.liferay.com/geolite.maxmind.com/download/geoip/database/GeoLiteCity.dat.xz
gunzip GeoIP.dat.gz
unxz GeoLiteCity.dat.xz
```
