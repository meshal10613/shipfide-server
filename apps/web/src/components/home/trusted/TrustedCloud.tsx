'use client';

import { motion } from 'framer-motion';

const PARTNERS = [
  'TechStore BD',
  'Glamour Closet',
  'Apex Electronics',
  'Aarong Craft',
  'Daraz Partner',
  'Chaldal Express',
];

export const TrustedCloud = () => {
  return (
    <section className="py-12 bg-white border-b border-slate-200/60">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6">
        <p className="text-xs font-bold uppercase tracking-widest text-slate-400">
          Trusted by 12,500+ High-Volume E-Commerce Brands & Enterprises
        </p>

        <div className="flex flex-wrap items-center justify-center gap-8 sm:gap-12 md:gap-16 opacity-75 grayscale hover:grayscale-0 transition-all">
          {PARTNERS.map((partner, index) => (
            <motion.div
              key={partner}
              initial={{ opacity: 0, y: 10 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true }}
              transition={{ duration: 0.4, delay: index * 0.1 }}
              className="text-sm sm:text-base font-black text-[#404042] tracking-wider hover:text-[#DD0033] transition-colors cursor-pointer"
            >
              {partner}
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  );
};
