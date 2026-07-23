'use client';

import { motion } from 'framer-motion';
import { SectionHeading } from '@/components/common/SectionHeading';
import { FEATURES_DATA } from '@/constants/landing';
import {
  Navigation,
  Truck,
  Building2,
  BarChart3,
  ShieldAlert,
  Wallet,
  Code2,
  FileSpreadsheet,
  Zap,
} from 'lucide-react';

const ICON_MAP: Record<string, React.ElementType> = {
  Navigation,
  Truck,
  Building2,
  BarChart3,
  ShieldAlert,
  Wallet,
  Code2,
  FileSpreadsheet,
  Zap,
};

export const FeaturesSection = () => {
  return (
    <section className="py-20 bg-[#F9FCFE] border-b border-slate-200/60" id="features">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-16">
        <SectionHeading
          badge="ENTERPRISE CAPABILITIES"
          title="Everything You Need to Scale Logistics Nationwide"
          description="Powered by cutting-edge telemetry, automated hub routing, and instant COD treasury management."
        />

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 sm:gap-8">
          {FEATURES_DATA.map((feature, index) => {
            const IconComponent = ICON_MAP[feature.iconName] || Navigation;

            return (
              <motion.div
                key={feature.id}
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                viewport={{ once: true }}
                transition={{ duration: 0.5, delay: index * 0.08 }}
                whileHover={{ y: -5 }}
                className="bg-white p-7 rounded-3xl border border-slate-200 shadow-sm hover:shadow-xl hover:border-[#DD0033]/30 transition-all flex flex-col justify-between group relative overflow-hidden"
              >
                {/* Accent Corner Bar */}
                <div className="absolute top-0 right-0 w-24 h-24 bg-gradient-to-bl from-[#DD0033]/5 to-transparent rounded-bl-full pointer-events-none group-hover:from-[#DD0033]/10 transition-colors" />

                <div className="space-y-4 relative z-10">
                  <div className="flex items-center justify-between">
                    <div className="w-12 h-12 rounded-2xl bg-[#DD0033]/10 text-[#DD0033] flex items-center justify-center group-hover:bg-[#DD0033] group-hover:text-white transition-colors shadow-xs">
                      <IconComponent className="w-6 h-6" />
                    </div>
                    {feature.badge && (
                      <span className="px-2.5 py-0.5 rounded-full text-[10px] font-bold uppercase tracking-wider bg-slate-100 text-[#404042] border border-slate-200">
                        {feature.badge}
                      </span>
                    )}
                  </div>

                  <h3 className="text-xl font-extrabold text-[#404042] tracking-tight group-hover:text-[#DD0033] transition-colors">
                    {feature.title}
                  </h3>

                  <p className="text-xs sm:text-sm text-slate-600 leading-relaxed">
                    {feature.description}
                  </p>
                </div>
              </motion.div>
            );
          })}
        </div>
      </div>
    </section>
  );
};
