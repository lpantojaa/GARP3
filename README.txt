-populationapi.go
Es el programa que hace la conexión con la base de datos y recibe las peticiones del asistente en formato JSON.

-Ejemplo JSON de la petición a la BD.txt
Es un ejemplo de lo que devuelve el asistente através de Actions On Google en formato JSON. Incluye las 3 variables, país, año y grupo de edad.

-ngrok.exe
Es el puente entre Actions On Google y populatioapi.go. Utiliza protocolo http y se debe asignar un puerto de entrada.