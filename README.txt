-populationapi.go
Es el programa que hace la conexi�n con la base de datos y recibe las peticiones del asistente en formato JSON.

-Ejemplo JSON de la petici�n a la BD.txt
Es un ejemplo de lo que devuelve el asistente atrav�s de Actions On Google en formato JSON. Incluye las 3 variables, pa�s, a�o y grupo de edad.

-ngrok.exe
Es el puente entre Actions On Google y populatioapi.go. Utiliza protocolo http y se debe asignar un puerto de entrada.