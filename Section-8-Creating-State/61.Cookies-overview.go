// Here we will learn, how to write cookies to the client's browser.
// Also how to read cookie that are sent from the client to the server.
// The way cookies work is that cookie is a litle file that can hold information.
// Server can write cookie to the client's machine if the client allows for the cookies to be written.
// If someone doesn't allow for the cookies to be written the workaround is creating state.
// How is state created from cookies.
// If we are able to write cookie to a client's machine. Anytime client makes a request back to the server,
// the client will send that cookie.
// We (server) can put into that cookie a unique id. Then we (server) can say, do you have a cookie and what is the
// unique id in it.
// Then we (server) can look at that unique id and can tell what user is it's associated with.
// and that user needs to log back in since it's been too long since they have been here
// OR they are already authenticated and can see this other information they are requesting and
// we (server) will send that information back to them.
// This allows us to uniquely identify each user. By putting a unique id in the cookie
// and that cookie gets sent back to our machine via a request.
// A domain can write a cookie in our browser and anytime the browser makes request back
// to the domain, the browser will look and see if there's a cookie for that domain, if so
// then the browser will send that cookie back to the domain. Cookies are domain specific.
// Say if pepsi writes a cookie and coke writes a cookie, then the browser will only send
// cookie for pepsi back when the request is made to pepsi.
// Cookies are Domain specific.
// We can write a unique id to the cookie and that's how track the state of the application is
// for that one user.
// If the user doesn't accept cookie we can also pass a unique id through the url.
// Any link on the page we can add the unique id into the link, and then we just always use
// HTTPS and everything from the client to the server is encrypted. So it's whether a cookie
// being passed or some unique identifier in the url,
// all of it is encrypted as it is passed in the open feilds of the internet.
// On client the information is decrypted and then again encrypted when it is sent back out to the server.
// Unique id wil show up decrypted in the client's browser and the client's url bar.
// so if someone take a pic of the url bar then they could hijack the user's session by entering the information
// into the browser.
// That is one security hole in putting unique id in all of the links/urls in the broser page.
// That is why we learnt how to pass in the values through the URL.
