# Seminario_Go
 ## TP Entregagable realizado en el seminario de Golang TUDAI 2021

 Se crea una función que dada una cadena nde caracteres con un formato predeterminado genera una instancia de una estructura con los valores de los campos correspondientes.
 Para ello, al ingresar una cadena de caracteres la función generarResultado(cadena string), genera un objeto del tipo CadenaEntante que contiene un campo Value con el string ingresado.
 A partir de ese objeto se parsea la cadena mediante las funciones getTipo(cadena.Value) para obtener los 2 primeros caracteres que corresponden al tipo; getValue(cadena.Value) para obtener el valor correspondiente, getLength(cadena.Value) para obtener el largo de dicho valor.
 Luego se realizan chequeos para corroborar que la cadena sea válida según las especificaciones de la consigna, para finalizar si el mismo es exitoso se crea la estructura pedida con tipo, largo y valor,
 de lo contrario imprime un error.
 Se generaron chequeos, main_test.go y entities_test.go para las funciones implementadas asi como también outTesting y outTesting2 con sus respectivos html para comprobar el testeo en cadenas válidas e inválidas.


 Integrantes: Maria Elva Kehler Diviasi - Sergio Marcelo Yañez 
        DNI :      24914542                     23322529
