importScripts('https://www.gstatic.com/firebasejs/7.19.1/firebase-app.js');
importScripts('https://www.gstatic.com/firebasejs/7.19.1/firebase-messaging.js');

firebase.initializeApp({
    apiKey: "AIzaSyAkMc-_O5v5PmLVGVZP-Oaw2T4yLiCfWm0",
    authDomain: "tavanbogd-44377.firebaseapp.com",
    projectId: "tavanbogd-44377",
    storageBucket: "tavanbogd-44377.appspot.com",
    messagingSenderId: "982500571278",
    appId: "1:982500571278:web:249fbed823d74fcf7231df",
    measurementId: "G-493G7B3240"
});

const messaging = firebase.messaging();

messaging.setBackgroundMessageHandler(function(payload) {
    console.log('[firebase-messaging-sw.js] Received background message ', payload);
    // Customize notification here
    const notificationTitle = 'Background Message Title';
    const notificationOptions = {
        body: 'Background Message body.',
        icon: '/firebase-logo.png'
    };

    return self.registration.showNotification(notificationTitle,
        notificationOptions);
});
