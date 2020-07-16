# WebsiteStatusTracker 
We will be building a simple website status checker application using concurrency features of Golang. In this application, we will have a collection (slice) of websites and we will be running a loop on this given slice of websites. In each iteration, we will be calling getWebsite function to perform a get request on that given website.

## Branch

- **WST1** -> contains waitgroup implementaion

- **WST2** -> contains waitgroup+mutex implementation

- **WST3** -> contains channel implementaion

Medium Link https://medium.com/@dipesh.kc/practical-example-of-concurrency-on-golang-fc4609ea8ed1
