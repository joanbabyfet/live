const videoElement = document.getElementById('video')

if (flvjs.isSupported()) {

    const flvPlayer = flvjs.createPlayer({

        type: 'flv',

        url: 'http://127.0.0.1:8080/live/live_1003.flv'

    })

    flvPlayer.attachMediaElement(videoElement)

    flvPlayer.load()

    flvPlayer.play()

} else {

    console.log('FLV not supported')
}