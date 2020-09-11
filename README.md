## MYSQL-HEALTHCHECK 

This tool is util to run it on a mysql server as a healthcheck when you are using an AWS ELB. 


### Steps
* build 

go get -u github.com/go-sql-driver/mysql
go get -u  gopkg.in/gcfg.v1
go build

* Crear carpeta /opt/mysql-healthcheck y copiar archivo mysql-healthcheck
* copiar archivo mysql-healthcheck.service en /etc/init.d/mysql-healthcheck.service
* Dar permisos: chmod +x mysql-healthcheck.service
* crear archivo /etc/.mysql_healthcheck.conf con  datos de acceso a mysql
* exportar variable MYSQL_HEALTHCHECK_PATH  con el path del archivo /etc/.mysql_healthcheck.conf
