<script setup>
import { ArrowRight } from '@lucide/vue'
import CampaignCard from './CampaignCard.vue'

const campaigns = [
  {
    id: 1,
    title: "Amara Vegetable Farm Expansion",
    description: "Help us buy new irrigation equipment to increase our yield and supply more fresh vegetables to the local market.",
    image: "photo-1592982537447-7440770cbfc9",
    category: "AGRICULTURE",
    daysLeft: 12,
    avatar: "A",
    avatarColor: "#2A5F45",
    location: "Nairobi, Kenya",
    raised: 1240,
    goal: 2000,
    backers: 34
  },
  {
    id: 2,
    title: "Tech Hub Community Center",
    description: "Building a space for local youth to learn programming and digital skills for the modern economy.",
    image: "photo-1517048676732-d65bc937f952",
    category: "EDUCATION",
    daysLeft: 5,
    avatar: "T",
    avatarColor: "#D4891A",
    location: "Lagos, Nigeria",
    raised: 4500,
    goal: 5000,
    backers: 112
  },
  {
    id: 3,
    title: "Solar Panels for Clinic",
    description: "Providing reliable electricity for the rural maternal health clinic to operate safely at night.",
    image: "photo-1509391366360-1161d6fb35db",
    category: "HEALTHCARE",
    daysLeft: 21,
    avatar: "M",
    avatarColor: "#1a4789",
    location: "Kigali, Rwanda",
    raised: 890,
    goal: 3500,
    backers: 18
  }
]
</script>

<template>
  <section id="campaigns" class="py-20 md:py-28">
    <div class="max-w-6xl mx-auto px-5">

      <div class="flex flex-col md:flex-row md:items-end justify-between mb-12 gap-4">
        <div>
          <p class="text-xs font-bold tracking-widest text-accent uppercase mb-3" style="font-family: 'DM Mono', monospace;">
            Live now
          </p>
          <h2 class="text-3xl md:text-4xl font-bold text-foreground" style="font-family: 'Fraunces', serif;">
            Active campaigns
          </h2>
        </div>
        <button class="inline-flex items-center gap-2 text-sm font-semibold text-primary hover:underline underline-offset-4 transition-all" style="font-family: 'DM Sans', sans-serif;">
          See all campaigns
          <ArrowRight :size="15" />
        </button>
      </div>

      <div class="grid md:grid-cols-3 gap-6">
        <CampaignCard
            v-for="c in campaigns"
            :key="c.id"
            :campaign="c"
        />
      </div>

    </div>
  </section>
</template>

