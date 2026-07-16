<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { Send } from '@lucide/vue'

const stats = [
  { value: "48K+", label: "Anggota" },
  { value: "$2.4M", label: "Terkumpul" },
  { value: "94%", label: "Tingkat pelunasan" },
]

const carouselImages = [
  "photo-1529156069898-49953e39b3ac",
  "photo-1531123897727-8f129e1688ce",
  "photo-1506277886164-e25aa3f4ef7f"
]

const currentImageIndex = ref(0)
let intervalId = null

onMounted(() => {
  intervalId = setInterval(() => {
    currentImageIndex.value = (currentImageIndex.value + 1) % carouselImages.length
  }, 5000)
})

onUnmounted(() => {
  if (intervalId) clearInterval(intervalId)
})
</script>

<template>
  <section id="home" class="relative overflow-hidden pt-20 pb-24 md:pt-32 md:pb-40 min-h-[90vh] flex items-center">

    <div class="absolute inset-0 z-0 bg-background">
      <Transition
          enter-active-class="transition-opacity duration-1000 ease-in-out absolute inset-0"
          enter-from-class="opacity-0"
          enter-to-class="opacity-100"
          leave-active-class="transition-opacity duration-1000 ease-in-out absolute inset-0"
          leave-from-class="opacity-100"
          leave-to-class="opacity-0"
      >
        <img
            :key="currentImageIndex"
            :src="`https://images.unsplash.com/${carouselImages[currentImageIndex]}?w=1920&h=1080&fit=crop&auto=format`"
            alt="Latar belakang komunitas"
            class="w-full h-full object-cover"
        />
      </Transition>

      <!-- Pembaruan Gradien Overlay untuk Visibilitas Gambar Maksimal -->
      <div class="absolute inset-0 bg-background/70 md:bg-gradient-to-r md:from-background/95 md:via-background/50 md:to-transparent backdrop-blur-[1px]"></div>
    </div>

    <div class="max-w-6xl mx-auto px-5 relative z-10 w-full">
      <div class="grid md:grid-cols-2 gap-12 items-center">

        <div>
          <div class="inline-flex items-center gap-2 bg-primary/15 text-primary text-xs font-semibold px-4 py-2 rounded-full mb-6 border border-primary/30 backdrop-blur-md">
            <span class="w-1.5 h-1.5 rounded-full bg-primary animate-pulse"></span>
            Dipercaya oleh lebih dari 48.000 anggota aktif
          </div>

          <h1
              class="text-4xl md:text-5xl lg:text-6xl font-bold text-foreground leading-[1.1] mb-5"
              style="font-family: 'Fraunces', serif; font-weight: 700;"
          >
            Uang Anda,<br />
            <span class="text-primary italic">komunitas Anda,</span><br />
            masa depan Anda.
          </h1>

          <p class="text-base md:text-lg text-muted-foreground leading-relaxed mb-8 max-w-md" style="font-family: 'DM Sans', sans-serif;">
            Umoja menghadirkan pinjaman mikro, tabungan kelompok, dan urun dana (crowdfunding) untuk siapa saja yang memiliki ponsel — tanpa rekening bank, tanpa dokumen rumit, dan tanpa saldo minimum.
          </p>

          <div class="flex items-center gap-6 mt-6 pt-8 border-t border-border/50">
            <div v-for="stat in stats" :key="stat.label">
              <div class="text-2xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
                {{ stat.value }}
              </div>
              <div class="text-xs text-muted-foreground mt-0.5" style="font-family: 'DM Mono', monospace;">
                {{ stat.label }}
              </div>
            </div>
          </div>
        </div>

        <div class="relative w-full max-w-sm mx-auto md:ml-auto mt-10 md:mt-0">

          <div class="bg-card/80 backdrop-blur-xl border border-border/60 rounded-3xl p-6 shadow-2xl">
            <div class="flex items-center justify-between mb-5">
              <div>
                <p class="text-xs font-bold text-muted-foreground tracking-wider mb-1" style="font-family: 'DM Mono', monospace;">KAMPANYE AKTIF</p>
                <p class="font-bold text-foreground text-lg leading-tight" style="font-family: 'Fraunces', serif;">Perluasan Kebun Sayur</p>
              </div>
              <div class="bg-primary/10 text-primary text-xs font-bold px-3 py-1.5 rounded-full">62%</div>
            </div>

            <div class="w-full h-2.5 bg-muted/80 rounded-full overflow-hidden mb-4">
              <div class="h-full bg-primary rounded-full transition-all duration-1000" style="width: 62%;"></div>
            </div>

            <div class="flex justify-between items-center">
              <div>
                <span class="text-sm font-bold text-foreground block" style="font-family: 'DM Mono', monospace;">$1.240</span>
                <span class="text-xs text-muted-foreground font-medium">Terkumpul</span>
              </div>
              <div class="text-right">
                <span class="text-sm font-bold text-foreground block" style="font-family: 'DM Mono', monospace;">$2.000</span>
                <span class="text-xs text-muted-foreground font-medium">Target</span>
              </div>
            </div>
          </div>

          <div class="absolute -bottom-6 -left-4 md:-left-12 bg-accent text-white rounded-2xl px-5 py-4 shadow-xl hidden sm:block animate-bounce" style="animation-duration: 3s;">
            <div class="flex items-center gap-3">
              <div class="bg-white/20 p-2 rounded-full flex-shrink-0">
                <Send :size="16" class="text-white" />
              </div>
              <div>
                <p class="text-sm font-bold leading-tight">Kontribusi baru</p>
                <p class="text-xs opacity-90 mt-0.5">Grace mengirim $25 · baru saja</p>
              </div>
            </div>
          </div>

        </div>

      </div>
    </div>
  </section>
</template>