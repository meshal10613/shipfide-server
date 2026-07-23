'use client';

import { motion } from 'framer-motion';
import { STATS_DATA } from '@/constants/landing';

export const StatsSection = () => {
  return (
    <section className="py-16 bg-[#404042] text-white border-b border-slate-700/50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="grid grid-cols-2 md:grid-cols-4 gap-8 text-center divide-y md:divide-y-0 md:divide-x divide-slate-700/60">
          {STATS_DATA.map((st, index) => (
            <motion.div
              key={st.label}
              initial={{ opacity: 0, scale: 0.9 }}
              whileInView={{ opacity: 1, scale: 1 }}
              viewport={{ once: true }}
              transition={{ duration: 0.5, delay: index * 0.1 }}
              className="pt-6 md:pt-0 space-y-2"
            >
              <div className="text-3xl sm:text-4xl md:text-5xl font-black text-white tracking-tight">
                {st.prefix}
                <span className="text-[#DD0033]">
                  {st.value >= 1000000 ? `${(st.value / 1000000).toFixed(1)}M` : st.value.toLocaleString()}
                </span>
                {st.suffix}
              </div>
              <p className="text-xs sm:text-sm font-semibold text-slate-300">
                {st.label}
              </p>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  );
};
