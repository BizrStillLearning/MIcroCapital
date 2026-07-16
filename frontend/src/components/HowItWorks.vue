<script setup>
import { ref } from 'vue'
import { Smartphone, Users, Sprout, ChevronDown } from '@lucide/vue'

const openIndex = ref(null)

const toggleAccordion = (index) => {
  openIndex.value = openIndex.value === index ? null : index
}

const steps = [
  {
    title: "Create an account",
    description: "Sign up using just your phone number. No bank account or paperwork required to get started.",
    icon: Smartphone,
    details: "Unduh aplikasi Umoja atau akses melalui kode USSD di HP Anda. Masukkan nomor telepon aktif untuk menerima kode OTP. Setelah terverifikasi, buat PIN 4-digit yang aman. Akun Anda langsung aktif dalam hitungan detik tanpa perlu dokumen KTP atau ke kantor cabang bank."
  },
  {
    title: "Join a community",
    description: "Connect with local groups or create your own savings circle with trusted friends and family.",
    icon: Users,
    details: "Cari grup komunitas terdekat atau mulai lingkaran tabungan (savings circle) Anda sendiri. Anda dapat mengundang keluarga atau tetangga hanya menggunakan nomor HP mereka. Tentukan aturan iuran dan target bersama secara transparan di dalam sistem."
  },
  {
    title: "Grow your funds",
    description: "Start contributing, applying for micro-loans, or raising money for your business needs.",
    icon: Sprout,
    details: "Mulai kumpulkan dana melalui iuran rutin. Saat Anda butuh modal, ajukan pinjaman mikro yang disetujui langsung oleh komunitas Anda. Anda juga bisa membuka kampanye crowdfunding publik agar lebih banyak orang bisa membantu mendanai usaha Anda."
  }
]
</script>

<template>
  <section id="how-it-works" class="py-20 md:py-28 bg-card border-y border-border">
    <div class="max-w-6xl mx-auto px-5">

      <div class="text-center mb-14">
        <p class="text-xs font-bold tracking-widest text-accent uppercase mb-3" style="font-family: 'DM Mono', monospace;">
          Simple process
        </p>
        <h2 class="text-3xl md:text-4xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
          Three steps to your first fund
        </h2>
      </div>

      <div class="grid md:grid-cols-3 gap-8 relative items-start">
        <div class="hidden md:block absolute top-10 left-1/4 right-1/4 h-px bg-gradient-to-r from-transparent via-border to-transparent"></div>

        <div v-for="(step, i) in steps" :key="i" class="relative text-center group flex flex-col items-center">

          <div class="flex justify-center mb-5">
            <div class="relative">
              <div class="w-20 h-20 rounded-2xl bg-primary/10 border-2 border-primary/20 flex items-center justify-center group-hover:bg-primary group-hover:border-primary transition-all duration-300">
                <component
                    :is="step.icon"
                    :size="32"
                    class="text-primary group-hover:text-primary-foreground transition-colors duration-300"
                />
              </div>

              <span
                  class="absolute -top-3 -right-3 w-7 h-7 rounded-full bg-accent text-white text-xs font-bold flex items-center justify-center"
                  style="font-family: 'DM Mono', monospace;"
              >
                {{ i + 1 }}
              </span>
            </div>
          </div>

          <h3 class="text-lg font-bold text-foreground mb-2" style="font-family: 'Fraunces', serif;">
            {{ step.title }}
          </h3>
          <p class="text-sm text-muted-foreground leading-relaxed mb-4" style="font-family: 'DM Sans', sans-serif;">
            {{ step.description }}
          </p>

          <button
              @click="toggleAccordion(i)"
              class="flex items-center justify-center gap-1.5 text-xs font-bold text-primary hover:text-accent transition-colors mt-auto"
              style="font-family: 'DM Mono', monospace;"
          >
            {{ openIndex === i ? 'Tutup' : 'Detail' }}
            <ChevronDown
                :size="14"
                class="transition-transform duration-300"
                :class="{ 'rotate-180': openIndex === i }"
            />
          </button>

          <div
              class="grid transition-all duration-300 ease-in-out w-full mt-3"
              :class="openIndex === i ? 'grid-rows-[1fr] opacity-100' : 'grid-rows-[0fr] opacity-0'"
          >
            <div class="overflow-hidden">
              <div class="bg-primary/5 border border-primary/10 rounded-xl p-4 text-left text-sm text-foreground leading-relaxed" style="font-family: 'DM Sans', sans-serif;">
                {{ step.details }}
              </div>
            </div>
          </div>

        </div>
      </div>

    </div>
  </section>
</template>