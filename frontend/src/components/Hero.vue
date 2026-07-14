<script setup>
import { ArrowRight, Send } from '@lucide/vue'

const stats = [
  { value: "48K+", label: "Members" },
  { value: "$2.4M", label: "Raised" },
  { value: "94%", label: "Repayment rate" },
]
</script>

<template>
  <section class="relative overflow-hidden pt-16 pb-20 md:pt-24 md:pb-32">
    <div class="absolute inset-0 pointer-events-none">
      <div
          class="absolute -top-32 -right-32 w-96 h-96 rounded-full opacity-20"
          style="background: radial-gradient(circle, #2A5F45, transparent);"
      ></div>
      <div
          class="absolute -bottom-20 -left-20 w-72 h-72 rounded-full opacity-15"
          style="background: radial-gradient(circle, #D4891A, transparent);"
      ></div>
    </div>

    <div class="max-w-6xl mx-auto px-5 relative">
      <div class="grid md:grid-cols-2 gap-12 items-center">

        <div>
          <div class="inline-flex items-center gap-2 bg-primary/10 text-primary text-xs font-semibold px-3 py-1.5 rounded-full mb-6 border border-primary/20">
            <span class="w-1.5 h-1.5 rounded-full bg-primary animate-pulse"></span>
            Trusted by 48,000+ members across Africa
          </div>

          <h1
              class="text-4xl md:text-5xl lg:text-6xl font-bold text-foreground leading-[1.1] mb-5"
              style="font-family: 'Fraunces', serif; font-weight: 700;"
          >
            Your money,<br />
            <span class="text-primary italic">your community,</span><br />
            your future.
          </h1>

          <p class="text-base md:text-lg text-muted-foreground leading-relaxed mb-8 max-w-md" style="font-family: 'DM Sans', sans-serif;">
            Umoja brings micro-loans, group savings, and crowdfunding to anyone with a phone — no bank account, no paperwork, no minimum balance required.
          </p>

          <div class="flex flex-col sm:flex-row gap-3">
            <button class="flex items-center justify-center gap-2 bg-primary text-primary-foreground px-6 py-3.5 rounded-xl font-semibold text-base hover:opacity-90 transition-all hover:shadow-lg hover:shadow-primary/20">
              Start raising funds
              <ArrowRight :size="18" />
            </button>
            <button class="flex items-center justify-center gap-2 bg-card border border-border text-foreground px-6 py-3.5 rounded-xl font-semibold text-base hover:bg-muted transition-colors">
              Browse campaigns
            </button>
          </div>

          <div class="flex items-center gap-6 mt-10 pt-8 border-t border-border">
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

        <div class="relative">
          <div class="relative rounded-2xl overflow-hidden shadow-2xl" style="background: #2A5F45;">
            <img
                src="https://images.unsplash.com/photo-1529156069898-49953e39b3ac?w=600&h=500&fit=crop&auto=format"
                alt="Community members working together"
                class="w-full h-72 md:h-96 object-cover opacity-80 mix-blend-multiply"
            />
            <div class="absolute inset-0 bg-gradient-to-t from-primary/80 via-transparent to-transparent"></div>

            <div class="absolute bottom-0 left-0 right-0 p-5">
              <div class="bg-card/95 backdrop-blur-sm rounded-xl p-4 shadow-lg">
                <div class="flex items-center justify-between mb-3">
                  <div>
                    <p class="text-xs text-muted-foreground" style="font-family: 'DM Mono', monospace;">ACTIVE CAMPAIGN</p>
                    <p class="font-semibold text-foreground text-sm mt-0.5" style="font-family: 'Fraunces', serif;">Amara Vegetable Farm</p>
                  </div>
                  <div class="bg-primary/10 text-primary text-xs font-bold px-2.5 py-1 rounded-full">62%</div>
                </div>
                <div class="w-full h-2 bg-muted rounded-full overflow-hidden">
                  <div class="h-full bg-primary rounded-full transition-all" style="width: 62%;"></div>
                </div>
                <div class="flex justify-between mt-2">
                  <span class="text-xs text-muted-foreground" style="font-family: 'DM Mono', monospace;">$1,240 raised</span>
                  <span class="text-xs text-muted-foreground" style="font-family: 'DM Mono', monospace;">$2,000 goal</span>
                </div>
              </div>
            </div>
          </div>

          <div class="absolute -top-4 -right-4 bg-accent text-white rounded-2xl px-4 py-3 shadow-lg hidden md:block">
            <div class="flex items-center gap-2">
              <Send :size="14" class="text-white" />
              <div>
                <p class="text-xs font-semibold">New contribution</p>
                <p class="text-xs opacity-80">Grace sent $25 · just now</p>
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>
  </section>
</template>

