self.addEventListener('push', (e) => {
    const message = e.data ? e.data.text() : 'no payload'

    let n = self.registration.showNotification(message)

    e.waitUntil(n)
})