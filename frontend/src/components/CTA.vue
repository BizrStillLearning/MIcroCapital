<script setup>
import { Users, ArrowRight } from '@lucide/vue'
</script>

<template>
  <section class="py-20 md:py-28 bg-secondary border-y border-border">
    <div class="max-w-3xl mx-auto px-5 text-center">

      <div
          class="inline-flex items-center gap-2 text-xs font-bold text-primary bg-primary/10 border border-primary/20 px-3 py-1.5 rounded-full mb-6"
          style="font-family: 'DM Mono', monospace;"
      >
        <Users :size="13" />
        Join 48,000+ members across 14 countries
      </div>

      <h2
          class="text-4xl md:text-5xl font-bold text-foreground mb-5 leading-tight"
          style="font-family: 'Fraunces', serif;"
      >
        Your first step<br />
        <span class="italic text-primary">starts here.</span>
      </h2>

      <p
          class="text-base text-muted-foreground mb-8 leading-relaxed max-w-md mx-auto"
          style="font-family: 'DM Sans', sans-serif;"
      >
        No minimum deposit. No credit history required. Just a phone number and a dream.
      </p>

      <div class="flex flex-col sm:flex-row gap-3 justify-center">
        <button class="flex items-center justify-center gap-2 bg-primary text-primary-foreground px-8 py-4 rounded-xl font-semibold text-base hover:opacity-90 transition-all hover:shadow-xl hover:shadow-primary/25">
          Open free account
          <ArrowRight :size="18" />
        </button>
        <button class="flex items-center justify-center gap-2 bg-card border border-border text-foreground px-8 py-4 rounded-xl font-semibold text-base hover:bg-muted transition-colors">
          Download the app
        </button>
      </div>

      <p class="text-xs text-muted-foreground mt-5" style="font-family: 'DM Mono', monospace;">
        Free to join · Works on any smartphone · 12 languages
      </p>

    </div>
  </section>
</template>

