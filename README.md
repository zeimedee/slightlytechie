# SLIGHTLY TECHIE CODE CHALLENGE

### A post api where user can
1create a post
2.view a post
3.view all post
4.update a post
5.delete a post

### the app can bi started by running
`go run main.go`
###  in the terminal or command prompt

### the api will be available on
`http://localhost:4004/api/v1/`

### to create a post,make a post request to 
`http://localhost:4004/api/v1/add`

1.
### with a request body using the example below
`
{
    "title":"some title"
    "body":"some body"
}
`

2.
### to view all posts,make a get request to 
`http://localhost:4004/api/v1/readall`

3.
### to view a post,make a get request to 
`http://localhost:4004/api/v1/read/:id`

### include the post id in the url parameters

4.
### to update a post,make a patch request to 
`http://localhost:4004/api/v1/update/:id`


### include the post id in the url parameters

5.

### to delete a post,make a delete request to 
`http://localhost:4004/api/v1/delete/:id`

### include the post id in the url parameters

