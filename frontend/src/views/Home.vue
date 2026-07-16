<script setup>
import Navbar from '../components/Navbar.vue'
import Hero from '../components/Hero.vue'
import HowItWorks from '../components/HowItWorks.vue'
import Services from '../components/Services.vue'
import Trust from '../components/Trust.vue'
import Testimonials from '../components/Testimonials.vue'
import Footer from '../components/Footer.vue'

const handleNav = (section) => {
  const el = document.getElementById(section)
  if (el) {
    el.scrollIntoView({ behavior: 'smooth' })
  }
}
</script>

<template>
  <div class="min-h-screen bg-background" style="font-family: 'DM Sans', sans-serif;">
    <Navbar @nav="handleNav" />

    <main>
      <div id="hero"><Hero /></div>
      <div id="services"><Services /></div>
      <div id="how-it-works"><HowItWorks /></div>
      <Trust />
      <div id="community"><Testimonials /></div>
    </main>

    <Footer />
  </div>
</template>