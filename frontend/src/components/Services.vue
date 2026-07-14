<script setup>
import {
  ArrowRight,
  CircleDollarSign,
  Users,
  Wallet,
  WifiOff
} from '@lucide/vue'

const features = [
  {
    label: "Micro-loans",
    desc: "Instant approval up to $50 without credit history.",
    icon: CircleDollarSign
  },
  {
    label: "Group Savings",
    desc: "Pool funds securely with your local community.",
    icon: Users
  },
  {
    label: "Digital Wallet",
    desc: "Store and manage your funds safely in one place.",
    icon: Wallet
  },
  {
    label: "Offline Mode",
    desc: "Access services via USSD without internet.",
    icon: WifiOff
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

    <div class="max-w-6xl mx-auto px-5 relative">
      <div class="grid md:grid-cols-2 gap-12 items-center">

        <div>
          <p class="text-xs font-bold tracking-widest text-accent uppercase mb-4" style="font-family: 'DM Mono', monospace;">
            Financial tools
          </p>
          <h2 class="text-3xl md:text-4xl font-bold mb-5" style="font-family: 'Fraunces', serif;">
            Everything you need,<br />
            <span class="italic opacity-80">nothing you don't.</span>
          </h2>
          <p class="text-base opacity-75 leading-relaxed mb-8" style="font-family: 'DM Sans', sans-serif;">
            Built for 2G networks and low-end phones. Our tools are designed for real-world communities — fast, offline-first, and available in 12 local languages.
          </p>
          <button class="inline-flex items-center gap-2 bg-accent text-white px-6 py-3.5 rounded-xl font-semibold hover:opacity-90 transition-opacity">
            Explore all services
            <ArrowRight :size="18" />
          </button>
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div
              v-for="f in features"
              :key="f.label"
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
  </section>
</template>

