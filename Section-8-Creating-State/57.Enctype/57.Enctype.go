//We saw in the previous examples that we were uploading a file
// We created a form on a webpage
// And the user uploaded the file. That file got uploaded on the server.

// He pulled a different code, but we did this in 56-1 and 56-2 files.
// <form method ="POST" enctype="multipart/form-data">
// This enctype here allowed us to upload the file.
// So anytime you are going to create a form that allows user to upload a file, we have to
// set the enctype to multipart/form-data.
// enctype is the part of the "html form"
// default value of enctype is "application/x-www-form-urlencoded"
// so when we have <form method ="POST"> that is the default enctype.
// "urlencoded" gives us key value pairs seperated by '&' just as if they were in the URL.
// Remmeber it's the method that determines that wether it's going through the body or the URL.
// another enctype is "text/plain"
// <form method = "POST" enctype="text/plain">
// we only wan tto use text/plain for debugging.