importScripts('https://www.gstatic.com/firebasejs/7.19.1/firebase-app.js');
importScripts('https://www.gstatic.com/firebasejs/7.19.1/firebase-messaging.js');

firebase.initializeApp({
    apiKey: "AIzaSyAxLGoG4teww5fIaNTJ3t2Y70k2FCO8O68",
    authDomain: "b2b-finance.firebaseapp.com",
    projectId: "b2b-finance",
    storageBucket: "b2b-finance.appspot.com",
    messagingSenderId: "930334511447",
    appId: "1:930334511447:web:64cc282890d29b502eee12",
    measurementId: "G-YHQB6W0029",
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
