'use client';

import { AnnouncementBar } from './announcement/AnnouncementBar';
import { HeroSection } from './hero/HeroSection';
import { TrustedCloud } from './trusted/TrustedCloud';
import { FeaturesSection } from './features/FeaturesSection';
import { WhyChooseSection } from './why-choose/WhyChooseSection';
import { ServicesSection } from './services/ServicesSection';
import { TrackingPreview } from './tracking/TrackingPreview';
import { DashboardPreview } from './dashboard-preview/DashboardPreview';
import { HowItWorks } from './how-it-works/HowItWorks';
import { TestimonialsSection } from './testimonials/TestimonialsSection';
import { StatsSection } from './stats/StatsSection';
import { FaqSection } from './faq/FaqSection';
import { CtaSection } from './cta/CtaSection';

export const LandingPage = () => {
  return (
    <div className="bg-[#F9FCFE] text-[#404042] font-sans antialiased">
      {/* 1. Announcement Bar */}
      <AnnouncementBar />

      {/* 2. Hero Section */}
      <HeroSection />

      {/* 3. Trusted Companies Cloud */}
      <TrustedCloud />

      {/* 4. Enterprise Features Grid */}
      <FeaturesSection />

      {/* 5. Why Choose Shipfide */}
      <WhyChooseSection />

      {/* 6. Shipping Services */}
      <ServicesSection />

      {/* 7. Live Tracking Telemetry Preview */}
      <TrackingPreview />

      {/* 8. Executive Dashboard Preview */}
      <DashboardPreview />

      {/* 9. How It Works (Workflow Stepper) */}
      <HowItWorks />

      {/* 10. Customer Testimonials */}
      <TestimonialsSection />

      {/* 11. Animated Statistics Counters */}
      <StatsSection />

      {/* 12. Frequently Asked Questions (FAQ) */}
      <FaqSection />

      {/* 13. Call To Action (CTA) */}
      <CtaSection />
    </div>
  );
};
