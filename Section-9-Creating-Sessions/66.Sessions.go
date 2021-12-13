// So far we have seen how to craetea a cookie, and pass data through a form.
// We want to give user a unique identifier/ unique number and we will store that in a cookie.
// When a user makes a request to the server they will sends us that unique id.
// We can then take that unique id and know who was communicating with the server.
// This allows us to maintain information on the server about certain individuals.
// Clienta make request tot he server and we want the clients to send in a unique identifier to the server
// and when we get that unique id we could associate that id with bunch of other information on the server that
// let's us know who it is and what access credentials and authority they have, how much is in their bank account etc.
// So we can store the unique id in the cookie. cookies are domain specific so, every request client makes
// to ur domain, if there is a cookie in the client for our domain, client will send that cookie to us
// we can grab that value and then uniquely identify the client.
// This is how we create state.
// even though the client and server relation is not on going, we can create this on going relationship by passing
// in this unique id back and forth.
// One place we store the unique id is in a cookie. If we wanted to we can store that unique id in the URL.
// So in every link on our page (as the user would lcick on different links to go on different sites.)
// appended to the links could be a url parameter (variable in the url)  with a unique identifier that will allow
// us to see who is the person even though they wouldn't have a cookie. This is how we could dynamically
// sticking the unique id in the each url and taking it our everytime the client makes a request,
// because that would be part of what they are sending to us in the url.
// If we are using HTTPS so that's all encrypted, so the urls are encrypted.
// Now we'll look at creating a session, like someone would log in they could do suff.
// With authorized access credentials to certain areas and then log out. We could create that state.