<script setup>
import { ref } from 'vue'
import {
  CircleDollarSign,
  Users,
  Wallet,
  WifiOff,
  X
} from '@lucide/vue'

const selectedFeature = ref(null)

const features = [
  {
    label: "Micro-loans",
    desc: "Persetujuan instan hingga $50 tanpa riwayat kredit.",
    icon: CircleDollarSign,
    detail: "Dapatkan persetujuan instan untuk pinjaman hingga $50 tanpa memerlukan riwayat kredit tradisional dari bank. Sistem kami menggunakan algoritma pintar berbasis komunitas untuk memvalidasi profil Anda. Dana akan langsung masuk ke dompet digital Anda dan dapat ditarik kapan saja."
  },
  {
    label: "Group Savings",
    desc: "Kumpulkan dana dengan aman bersama komunitas lokal.",
    icon: Users,
    detail: "Kumpulkan dana bersama komunitas atau keluarga secara aman. Anda dapat mengatur target tabungan, jadwal iuran, dan memantau kontribusi setiap anggota secara transparan. Sangat cocok untuk modal usaha bersama atau dana darurat komunitas."
  },
  {
    label: "Digital Wallet",
    desc: "Simpan dan kelola dana Anda dengan aman di satu tempat.",
    icon: Wallet,
    detail: "Simpan dan kelola uang Anda dalam satu dompet digital yang dilindungi oleh enkripsi tingkat bank (256-bit). Lakukan transfer ke sesama pengguna secara gratis, bayar tagihan, atau cairkan dana ke agen lokal terdekat tanpa biaya tersembunyi."
  },
  {
    label: "Offline Mode",
    desc: "Akses layanan via USSD tanpa koneksi internet.",
    icon: WifiOff,
    detail: "Tidak ada kuota internet? Tidak masalah. Akses seluruh layanan keuangan Umoja menggunakan kode panggilan USSD standar di HP jenis apa pun. Transaksi Anda tetap diproses secara real-time dan aman layaknya menggunakan aplikasi."
  }
]
</script>

<template>
  <section id="services" class="py-20 md:py-28 bg-primary text-primary-foreground relative overflow-hidden">

    <div class="absolute inset-0 pointer-events-none">
      <div
          class="absolute top-0 right-0 w-80 h-80 rounded-full opacity-10"
          style="background: radial-gradient(circle, #D4891A, transparent);"
      ></div>
      <div
          class="absolute bottom-0 left-0 w-64 h-64 rounded-full opacity-10"
          style="background: radial-gradient(circle, #F5F0E8, transparent);"
      ></div>
    </div>

    <div class="max-w-6xl mx-auto px-5 relative z-10">
      <div class="grid md:grid-cols-2 gap-12 items-center">

        <div>
          <p class="text-xs font-bold tracking-widest text-accent uppercase mb-4" style="font-family: 'DM Mono', monospace;">
            Fitur Keuangan
          </p>
          <h2 class="text-3xl md:text-4xl font-bold mb-5" style="font-family: 'Fraunces', serif;">
            Semua yang Anda butuhkan,<br />
            <span class="italic opacity-80">tanpa hal yang rumit.</span>
          </h2>
          <p class="text-base opacity-75 leading-relaxed" style="font-family: 'DM Sans', sans-serif;">
            Dibangun untuk jaringan 2G dan ponsel sederhana. Fitur kami dirancang untuk komunitas di dunia nyata — cepat, dapat digunakan tanpa internet (offline-first), dan tersedia dalam 12 bahasa lokal.
          </p>
        </div>

        <div class="grid grid-cols-2 gap-4 mt-8 md:mt-0">
          <div
              v-for="f in features"
              :key="f.label"
              @click="selectedFeature = f"
              class="bg-white/10 border border-white/20 rounded-2xl p-5 hover:bg-white/15 transition-colors cursor-pointer group"
          >
            <div class="w-12 h-12 rounded-xl bg-white/15 flex items-center justify-center mb-4 group-hover:bg-accent transition-colors">
              <component :is="f.icon" :size="22" class="text-white" />
            </div>
            <h3 class="font-bold text-white mb-1" style="font-family: 'Fraunces', serif;">
              {{ f.label }}
            </h3>
            <p class="text-xs opacity-65" style="font-family: 'DM Sans', sans-serif;">
              {{ f.desc }}
            </p>
          </div>
        </div>

      </div>
    </div>

    <Transition
        enter-active-class="transition-opacity duration-500"
        enter-from-class="opacity-0"
        enter-to-class="opacity-100"
        leave-active-class="transition-opacity duration-300"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
    >
      <div
          v-if="selectedFeature"
          class="fixed inset-0 z-[60] bg-black/60 backdrop-blur-sm"
          @click="selectedFeature = null"
      ></div>
    </Transition>

    <Transition
        enter-active-class="transition-all duration-700 ease-[cubic-bezier(0.34,1.56,0.64,1)]"
        enter-from-class="opacity-0 scale-50 rotate-[-120deg]"
        enter-to-class="opacity-100 scale-100 rotate-0"
        leave-active-class="transition-all duration-500 ease-in-out"
        leave-from-class="opacity-100 scale-100 rotate-0"
        leave-to-class="opacity-0 scale-50 rotate-[120deg]"
    >
      <div
          v-if="selectedFeature"
          class="fixed inset-0 z-[70] flex items-center justify-center p-5 pointer-events-none"
      >
        <div class="bg-card border border-border rounded-3xl p-6 md:p-8 max-w-md w-full shadow-2xl pointer-events-auto relative">

          <button
              @click="selectedFeature = null"
              class="absolute top-5 right-5 text-muted-foreground hover:text-foreground bg-muted hover:bg-border p-2 rounded-full transition-colors"
          >
            <X :size="16" />
          </button>

          <div class="w-16 h-16 rounded-2xl bg-primary/10 border-2 border-primary/20 flex items-center justify-center mb-6">
            <component :is="selectedFeature.icon" :size="32" class="text-primary" />
          </div>

          <h3 class="text-2xl font-bold text-foreground mb-2" style="font-family: 'Fraunces', serif;">
            {{ selectedFeature.label }}
          </h3>

          <div class="inline-block bg-accent/10 text-accent text-xs font-bold px-3 py-1.5 rounded-full mb-5" style="font-family: 'DM Mono', monospace;">
            Fitur Utama Umoja
          </div>

          <p class="text-sm md:text-base text-muted-foreground leading-relaxed" style="font-family: 'DM Sans', sans-serif;">
            {{ selectedFeature.detail }}
          </p>

        </div>
      </div>
    </Transition>

  </section>
</template>