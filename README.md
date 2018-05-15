# go-webots

## config/database.json
<pre>
{
    "Port"     : 3306,
    "Host"     : "127.0.0.1",
    "Dbname"   : "test_db",
    "User"     : "user1",
    "Password" : "pass123",
    "Charset"  : "utf8mb4"
}
</pre>

## Download Website
/usr/bin/wget --mirror --convert-links --backup-converted -w 2 -np -A *.html,*.htm -R  jpg,jpeg,png,gif,tif,pdf,ppt,css,js \
  --referer="" --user-agent="Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.1) Gecko/20061204 Firefox/2.0.0.1" \
  http://www.example.com/ -o ./logs/w1.log
