<script setup>
import { Smartphone, Users, Sprout } from '@lucide/vue'

const steps = [
  {
    title: "Create an account",
    description: "Sign up using just your phone number. No bank account or paperwork required to get started.",
    icon: Smartphone
  },
  {
    title: "Join a community",
    description: "Connect with local groups or create your own savings circle with trusted friends and family.",
    icon: Users
  },
  {
    title: "Grow your funds",
    description: "Start contributing, applying for micro-loans, or raising money for your business needs.",
    icon: Sprout
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

      <div class="grid md:grid-cols-3 gap-8 relative">
        <div class="hidden md:block absolute top-10 left-1/4 right-1/4 h-px bg-gradient-to-r from-transparent via-border to-transparent"></div>

        <div v-for="(step, i) in steps" :key="i" class="relative text-center group">
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
          <p class="text-sm text-muted-foreground leading-relaxed" style="font-family: 'DM Sans', sans-serif;">
            {{ step.description }}
          </p>
        </div>

      </div>
    </div>
  </section>
</template>

