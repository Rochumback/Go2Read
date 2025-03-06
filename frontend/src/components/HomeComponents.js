import React, { useState, useEffect } from "react";
import Slider from "react-slick";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import "../styles/HomeStyles.css";

function MangaCarousel() {

    const [images, setImages] = useState([]);

    //useEffect(_ => {
    //  fetch("http://localhost:8080/api/images", {
    //      mode: "cors", headers: {
    //          'Access-Control-Allow-Origin': 'http://localhost:3000',
    //          'Access-Control-Allow-Credentials': true,
    //      }
    //  }).then(response => response.json())
    //      .then(json => {
    //          for (const image of json.images) {
    //              images.push((<div><img src={`${image}`} /></div>));
    //          }
    //      })
    //}, [])

    var settings = {
        dots: true,
        infinite: true,
        speed: 1000,
        slidesToShow: 1,
        slidesToScroll: 1,
        autoplaySpeed: 5000,
        autoplay: true,
    };

    return (
        <div class="slick-carousel">
            <Slider {...settings} style={{ "margin-left": 215 }}>
                {console.log(images)}
                {images}
            </Slider >
        </div>
    );
}

export { MangaCarousel }
