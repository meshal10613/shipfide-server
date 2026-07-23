'use client';

import { motion } from 'framer-motion';
import { SectionHeading } from '@/components/common/SectionHeading';
import { WHY_CHOOSE_DATA } from '@/constants/landing';
import { ShieldCheck, UserCheck, DollarSign, CheckCircle2 } from 'lucide-react';

export const WhyChooseSection = () => {
  return (
    <section className="py-20 bg-white border-b border-slate-200/60" id="why-choose">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-16">
        <div className="grid grid-cols-1 lg:grid-cols-12 gap-12 items-center">
          {/* Left Column: Copy & Points */}
          <motion.div
            initial={{ opacity: 0, x: -20 }}
            whileInView={{ opacity: 1, x: 0 }}
            viewport={{ once: true }}
            transition={{ duration: 0.6 }}
            className="lg:col-span-7 space-y-8"
          >
            <SectionHeading
              align="left"
              badge={WHY_CHOOSE_DATA.badge}
              title={WHY_CHOOSE_DATA.title}
              description={WHY_CHOOSE_DATA.description}
            />

            <div className="space-y-4 pt-2">
              {WHY_CHOOSE_DATA.points.map((pt, index) => (
                <div
                  key={index}
                  className="flex items-start gap-4 p-4 rounded-2xl bg-[#F9FCFE] border border-slate-200/80 shadow-xs"
                >
                  <div className="w-8 h-8 rounded-xl bg-[#DD0033]/10 text-[#DD0033] flex items-center justify-center shrink-0 mt-0.5">
                    <CheckCircle2 className="w-4 h-4" />
                  </div>
                  <div>
                    <h4 className="text-sm font-bold text-[#404042]">{pt.title}</h4>
                    <p className="text-xs text-slate-500 mt-0.5 leading-relaxed">{pt.desc}</p>
                  </div>
                </div>
              ))}
            </div>
          </motion.div>

          {/* Right Column: Grid Statistics */}
          <motion.div
            initial={{ opacity: 0, x: 20 }}
            whileInView={{ opacity: 1, x: 0 }}
            viewport={{ once: true }}
            transition={{ duration: 0.6, delay: 0.2 }}
            className="lg:col-span-5 grid grid-cols-2 gap-4"
          >
            {WHY_CHOOSE_DATA.stats.map((st, i) => (
              <div
                key={i}
                className="bg-[#F9FCFE] p-6 rounded-3xl border border-slate-200 shadow-sm text-center space-y-2 hover:border-[#DD0033]/40 transition-colors"
              >
                <span className="text-2xl sm:text-3xl font-black text-[#DD0033] block">
                  {st.value}
                </span>
                <span className="text-xs font-bold text-[#404042] block">
                  {st.label}
                </span>
              </div>
            ))}
          </motion.div>
        </div>
      </div>
    </section>
  );
};
