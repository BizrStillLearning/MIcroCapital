<script setup>
import {
  CheckCircle,
  Lock,
  ShieldCheck,
  Eye,
  Headset
} from '@lucide/vue'

const trustPoints = [
  {
    label: "256-bit Encryption",
    desc: "Bank-level security protecting your personal and financial data.",
    icon: Lock
  },
  {
    label: "Fraud Protection",
    desc: "24/7 AI-driven monitoring to detect and prevent suspicious activity.",
    icon: ShieldCheck
  },
  {
    label: "Full Transparency",
    desc: "Zero hidden fees. You always see exactly where your money goes.",
    icon: Eye
  },
  {
    label: "24/7 Support",
    desc: "Our dedicated local support team is always ready to help you.",
    icon: Headset
  }
]

const securityBadges = ["GSMA Certified", "ISO 27001", "PCI DSS Level 1"]
</script>

<template>
  <section id="trust" class="py-20 md:py-28 bg-card border-y border-border">
    <div class="max-w-6xl mx-auto px-5">

      <div class="text-center mb-14">
        <p class="text-xs font-bold tracking-widest text-accent uppercase mb-3" style="font-family: 'DM Mono', monospace;">
          Security first
        </p>
        <h2 class="text-3xl md:text-4xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
          Safe as a vault,<br />
          <span class="italic text-primary">simple as a handshake.</span>
        </h2>
      </div>

      <div class="grid sm:grid-cols-2 md:grid-cols-4 gap-6">
        <div v-for="t in trustPoints" :key="t.label" class="text-center group">
          <div class="w-16 h-16 rounded-2xl bg-primary/10 border border-primary/20 flex items-center justify-center mx-auto mb-4 group-hover:bg-primary transition-colors duration-300">
            <component
                :is="t.icon"
                :size="26"
                class="text-primary group-hover:text-primary-foreground transition-colors duration-300"
            />
          </div>
          <h3 class="font-bold text-foreground mb-1.5" style="font-family: 'Fraunces', serif;">
            {{ t.label }}
          </h3>
          <p class="text-sm text-muted-foreground" style="font-family: 'DM Sans', sans-serif;">
            {{ t.desc }}
          </p>
        </div>
      </div>

      <div class="mt-14 bg-primary/5 border border-primary/15 rounded-2xl p-6 md:p-8 flex flex-col md:flex-row items-center gap-6">
        <div class="flex-1">
          <h3 class="text-xl font-bold text-foreground mb-2" style="font-family: 'Fraunces', serif;">
            Regulated and licensed
          </h3>
          <p class="text-sm text-muted-foreground" style="font-family: 'DM Sans', sans-serif;">
            Umoja operates under Central Bank of West Africa frameworks and is an approved partner of the GSMA Mobile Money program. Funds are insured up to $10,000.
          </p>
        </div>

        <div class="flex items-center gap-4 flex-wrap justify-center">
          <div
              v-for="badge in securityBadges"
              :key="badge"
              class="flex items-center gap-1.5 bg-card border border-border px-4 py-2.5 rounded-xl shadow-sm"
          >
            <CheckCircle :size="14" class="text-primary" />
            <span class="text-xs font-semibold text-foreground" style="font-family: 'DM Mono', monospace;">
              {{ badge }}
            </span>
          </div>
        </div>
      </div>

    </div>
  </section>
</template>


