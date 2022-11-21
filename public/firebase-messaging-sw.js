importScripts('https://www.gstatic.com/firebasejs/7.19.1/firebase-app.js');
importScripts('https://www.gstatic.com/firebasejs/7.19.1/firebase-messaging.js');

firebase.initializeApp({
    apiKey: "AIzaSyCAnFbVIVTC_s8fkWDSFOGLvA5hGgPK3w0",
    authDomain: "tavanbogd-mmk.firebaseapp.com",
    projectId: "tavanbogd-mmk",
    storageBucket: "tavanbogd-mmk.appspot.com",
    messagingSenderId: "1058813657869",
    appId: "1:1058813657869:web:d2db0d7b885ce1f4c5e644",
    measurementId: "G-BHEEP0MHC2"
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
