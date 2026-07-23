'use client';

import { motion } from 'framer-motion';
import { SectionHeading } from '@/components/common/SectionHeading';
import { TESTIMONIALS_DATA } from '@/constants/landing';
import { Star, Quote } from 'lucide-react';
import Image from 'next/image';

export const TestimonialsSection = () => {
  return (
    <section className="py-20 bg-[#F9FCFE] border-b border-slate-200/60" id="testimonials">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-16">
        <SectionHeading
          badge="MERCHANT SUCCESS"
          title="Loved by E-Commerce Founders & Enterprise Operations"
          description="Hear how Shipfide transformed logistics, cash flow, and delivery speed for leading Bangladeshi brands."
        />

        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          {TESTIMONIALS_DATA.map((t, index) => (
            <motion.div
              key={t.id}
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true }}
              transition={{ duration: 0.5, delay: index * 0.1 }}
              className="bg-white p-8 rounded-3xl border border-slate-200 shadow-sm hover:shadow-xl transition-all flex flex-col justify-between space-y-6 relative"
            >
              <div className="space-y-4">
                {/* Rating Stars */}
                <div className="flex items-center gap-1 text-amber-500">
                  {Array.from({ length: t.rating }).map((_, i) => (
                    <Star key={i} className="w-4 h-4 fill-amber-500" />
                  ))}
                </div>

                {/* Quote Text */}
                <p className="text-xs sm:text-sm text-slate-700 leading-relaxed italic">
                  "{t.quote}"
                </p>
              </div>

              {/* Author & Profile Footer */}
              <div className="flex items-center gap-3 pt-4 border-t border-slate-100">
                <div className="w-11 h-11 rounded-full bg-slate-100 relative overflow-hidden border border-slate-200 shrink-0">
                  <Image
                    src={t.avatar}
                    alt={t.name}
                    fill
                    className="object-cover"
                  />
                </div>
                <div>
                  <h4 className="text-xs font-bold text-[#404042]">{t.name}</h4>
                  <p className="text-[11px] text-slate-500">{t.role} • {t.company}</p>
                  <span className="text-[10px] font-mono font-bold text-[#DD0033] mt-0.5 block">
                    {t.shipmentCount}
                  </span>
                </div>
              </div>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  );
};
