// pages/index.js
import React from "react";
import Hero from "../components/hero";
import Destination from "../components/destination";
import About from "../components/about";
import Review from "../components/review";
import Benefit from "../components/benefit";
import AOS from "aos";
import "aos/dist/aos.css";
import { useEffect } from "react";

function Home() {
  useEffect(() => {
    AOS.init({
      duration: 1500,
    });
  }, []);
  return (
    <div>
      <Hero />
      <Destination />
      <About />
      <Review />
      <Benefit />
    </div>
  );
}

export default Home;
