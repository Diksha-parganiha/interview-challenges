# League Backend Challenge

In main.go you will find a basic web server written in GoLang. It accepts a single request _/echo_. Extend the webservice with the ability to perform the following operations

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

1. Echo (given)
    - Return the matrix as a string in matrix format.
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.  

Run web server
```
go run .
```

Send request
```
$file = "C:\Users\aksha\Downloads\matrix.csv"
$url = "http://localhost:8080/echo"

# Create a boundary for multipart/form-data
$boundary = [System.Guid]::NewGuid().ToString()

# Build the multipart/form-data content
$multipartContent = @"
--$boundary
Content-Disposition: form-data; name="file"; filename="$(Split-Path $file -Leaf)"
Content-Type: application/octet-stream

$(Get-Content -Raw -Path $file)

--$boundary--
"@

# Convert the content to bytes
$bytes = [System.Text.Encoding]::UTF8.GetBytes($multipartContent)

# Make the web request
$response = Invoke-WebRequest -Uri $url -Method POST -ContentType "multipart/form-data; boundary=$boundary" -Body $bytes

# Output the response
$response

```
Invert request
```
$file = "C:\Users\aksha\Downloads\matrix.csv"
$url = "http://localhost:8080/invert"

# Create a boundary for multipart/form-data
$boundary = [System.Guid]::NewGuid().ToString()

# Build the multipart/form-data content
$multipartContent = @"
--$boundary
Content-Disposition: form-data; name="file"; filename="$(Split-Path $file -Leaf)"
Content-Type: application/octet-stream

$(Get-Content -Raw -Path $file)

--$boundary--
"@

# Convert the content to bytes
$bytes = [System.Text.Encoding]::UTF8.GetBytes($multipartContent)

# Make the web request
$response = Invoke-WebRequest -Uri $url -Method POST -ContentType "multipart/form-data; boundary=$boundary" -Body $bytes

# Output the response
$response

```
---Flatten Matrix
$file = "C:\Users\aksha\Downloads\matrix.csv"
$url = "http://localhost:8080/flatten"

# Create a boundary for multipart/form-data
$boundary = [System.Guid]::NewGuid().ToString()

# Build the multipart/form-data content
$multipartContent = @"
--$boundary
Content-Disposition: form-data; name="file"; filename="$(Split-Path $file -Leaf)"
Content-Type: application/octet-stream

$(Get-Content -Raw -Path $file)

--$boundary--
"@

# Convert the content to bytes
$bytes = [System.Text.Encoding]::UTF8.GetBytes($multipartContent)

# Make the web request
$response = Invoke-WebRequest -Uri $url -Method POST -ContentType "multipart/form-data; boundary=$boundary" -Body $bytes

# Output the response
$response

--- Sum
$url = "http://localhost:8080/sum"


## What we're looking for

- The solution runs
- The solution performs all cases correctly
- The code is easy to read
- The code is reasonably documented
- The code is tested
- The code is robust and handles invalid input and provides helpful error messages
