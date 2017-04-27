package models

func Setadm(adm string) string {
	var adserver string = `<!DOCTYPE html>
<html >
  <head>
    <meta charset="UTF-8">
    <title>adserver</title>
  </head>
  <body>`+adm+       
  `</body>
</html>`
    return adserver
}