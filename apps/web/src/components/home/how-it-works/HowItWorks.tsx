'use client';

import { motion } from 'framer-motion';
import { SectionHeading } from '@/components/common/SectionHeading';
import { HOW_IT_WORKS_STEPS } from '@/constants/landing';
import { PackagePlus, Truck, Building2, CheckCircle2 } from 'lucide-react';

const ICON_MAP: Record<string, React.ElementType> = {
  PackagePlus,
  Truck,
  Building2,
  CheckCircle2,
};

export const HowItWorks = () => {
  return (
    <section className="py-20 bg-white border-b border-slate-200/60" id="how-it-works">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-16">
        <SectionHeading
          badge="WORKFLOW"
          title="How Shipfide Delivers Your Parcels in 4 Simple Steps"
          description="Automated logistics pipeline connecting your warehouse directly to your receiver's doorstep."
        />

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8 relative">
          {HOW_IT_WORKS_STEPS.map((item, index) => {
            const IconComp = ICON_MAP[item.iconName] || CheckCircle2;

            return (
              <motion.div
                key={item.step}
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                viewport={{ once: true }}
                transition={{ duration: 0.5, delay: index * 0.12 }}
                className="bg-[#F9FCFE] p-7 rounded-3xl border border-slate-200 shadow-sm relative flex flex-col justify-between space-y-4 hover:border-[#DD0033]/40 transition-colors"
              >
                <div className="space-y-4">
                  <div className="flex items-center justify-between">
                    <span className="text-3xl font-black text-[#DD0033]/30 font-mono">
                      {item.step}
                    </span>
                    <div className="w-10 h-10 rounded-2xl bg-[#DD0033]/10 text-[#DD0033] flex items-center justify-center">
                      <IconComp className="w-5 h-5" />
                    </div>
                  </div>

                  <h3 className="text-lg font-extrabold text-[#404042]">
                    {item.title}
                  </h3>

                  <p className="text-xs text-slate-600 leading-relaxed">
                    {item.description}
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
