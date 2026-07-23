'use client';

import { useState } from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { SectionHeading } from '@/components/common/SectionHeading';
import { FAQ_DATA } from '@/constants/landing';
import { ChevronDown, HelpCircle } from 'lucide-react';

export const FaqSection = () => {
  const [openId, setOpenId] = useState<string | null>(FAQ_DATA[0].id);

  const toggleFaq = (id: string) => {
    setOpenId(openId === id ? null : id);
  };

  return (
    <section className="py-20 bg-[#F9FCFE] border-b border-slate-200/60" id="faq">
      <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 space-y-12">
        <SectionHeading
          badge="GOT QUESTIONS?"
          title="Frequently Asked Questions"
          description="Everything you need to know about COD disbursements, parcel security, coverage, and API integrations."
        />

        <div className="space-y-4">
          {FAQ_DATA.map((faq) => {
            const isOpen = openId === faq.id;

            return (
              <div
                key={faq.id}
                className="bg-white rounded-3xl border border-slate-200 overflow-hidden shadow-xs transition-colors"
              >
                <button
                  onClick={() => toggleFaq(faq.id)}
                  className="w-full p-6 text-left flex items-center justify-between gap-4 font-bold text-sm sm:text-base text-[#404042] hover:text-[#DD0033] transition-colors cursor-pointer"
                >
                  <div className="flex items-center gap-3">
                    <HelpCircle className="w-5 h-5 text-[#DD0033] shrink-0" />
                    <span>{faq.question}</span>
                  </div>
                  <ChevronDown
                    className={`w-5 h-5 text-slate-400 shrink-0 transition-transform duration-300 ${
                      isOpen ? 'rotate-180 text-[#DD0033]' : ''
                    }`}
                  />
                </button>

                <AnimatePresence>
                  {isOpen && (
                    <motion.div
                      initial={{ height: 0, opacity: 0 }}
                      animate={{ height: 'auto', opacity: 1 }}
                      exit={{ height: 0, opacity: 0 }}
                      transition={{ duration: 0.3 }}
                      className="overflow-hidden"
                    >
                      <div className="px-6 pb-6 pt-0 text-xs sm:text-sm text-slate-600 leading-relaxed border-t border-slate-100 mt-2">
                        <p className="pt-3">{faq.answer}</p>
                      </div>
                    </motion.div>
                  )}
                </AnimatePresence>
              </div>
            );
          })}
        </div>
      </div>
    </section>
  );
};
