<script setup>
const footerLinks = [
  {
    heading: "Platform",
    links: ["How it works", "Savings groups", "Micro loans", "Send money"],
  },
  {
    heading: "Support",
    links: ["Help center", "Community forum", "Safety & security", "Contact us", "Languages"],
  },
]

const legalLinks = ["Privacy", "Terms", "Cookies"]
</script>

<template>
  <footer class="bg-foreground text-primary-foreground py-14">
    <div class="max-w-6xl mx-auto px-5">
      <div class="grid md:grid-cols-4 gap-8 mb-10">

        <div class="md:col-span-2">
          <div class="flex items-center gap-2.5 mb-4">
            <div class="w-8 h-8 rounded-lg bg-primary flex items-center justify-center">
              <span class="text-white font-bold text-sm" style="font-family: 'Fraunces', serif;">U</span>
            </div>
            <span class="font-bold text-xl text-white" style="font-family: 'Fraunces', serif;">Umoja</span>
          </div>
          <p class="text-sm opacity-60 leading-relaxed max-w-xs" style="font-family: 'DM Sans', sans-serif;">
            Financial dignity for everyone, everywhere. Micro-capital crowdfunding for real communities.
          </p>
        </div>

        <div v-for="col in footerLinks" :key="col.heading">
          <h4 class="text-xs font-bold tracking-widest uppercase opacity-50 mb-4" style="font-family: 'DM Mono', monospace;">
            {{ col.heading }}
          </h4>
          <ul class="space-y-2.5">
            <li v-for="link in col.links" :key="link">
              <a href="#" class="text-sm opacity-60 hover:opacity-100 transition-opacity" style="font-family: 'DM Sans', sans-serif;">
                {{ link }}
              </a>
            </li>
          </ul>
        </div>

      </div>

      <div class="pt-8 border-t border-white/10 flex flex-col md:flex-row justify-between items-center gap-4">
        <p class="text-xs opacity-40" style="font-family: 'DM Mono', monospace;">
          © 2026 Umoja Financial. Licensed by CBWA · GSMA Partner
        </p>
        <div class="flex gap-5">
          <a
              v-for="item in legalLinks"
              :key="item"
              href="#"
              class="text-xs opacity-40 hover:opacity-70 transition-opacity"
              style="font-family: 'DM Mono', monospace;"
          >
            {{ item }}
          </a>
        </div>
      </div>

    </div>
  </footer>
</template>


