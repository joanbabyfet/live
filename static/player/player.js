const video = document.getElementById('video')

const videoSrc =
    'http://127.0.0.1:8080/live/live_1003.m3u8'

if (Hls.isSupported()) {

    const hls = new Hls()

    hls.loadSource(videoSrc)

    hls.attachMedia(video)

    hls.on(Hls.Events.MANIFEST_PARSED, function () {
        video.play()
    })
}