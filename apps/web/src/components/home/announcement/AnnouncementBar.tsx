'use client';

import { useState } from 'react';
import { ANNOUNCEMENT_DATA } from '@/constants/landing';
import { Sparkles, ArrowRight, X } from 'lucide-react';
import Link from 'next/link';

export const AnnouncementBar = () => {
  const [isVisible, setIsVisible] = useState(true);

  if (!isVisible) return null;

  return (
    <div className="bg-[#404042] text-white py-2.5 px-4 text-xs font-medium relative z-50 transition-all border-b border-slate-700/50">
      <div className="max-w-7xl mx-auto flex items-center justify-between gap-4">
        <div className="flex items-center gap-2.5 mx-auto text-center sm:text-left flex-wrap justify-center">
          <span className="inline-flex items-center gap-1.5 px-2.5 py-0.5 rounded-full text-[10px] font-black uppercase tracking-wider bg-[#DD0033] text-white">
            <Sparkles className="w-3 h-3" />
            {ANNOUNCEMENT_DATA.badge}
          </span>
          <span className="text-slate-200">{ANNOUNCEMENT_DATA.title}</span>
          <Link
            href={ANNOUNCEMENT_DATA.ctaUrl}
            className="inline-flex items-center gap-1 font-bold text-[#DD0033] hover:text-rose-400 hover:underline transition-colors ml-1"
          >
            {ANNOUNCEMENT_DATA.ctaText}
          </Link>
        </div>
        <button
          onClick={() => setIsVisible(false)}
          className="text-slate-400 hover:text-white p-1 rounded-full hover:bg-slate-700/50 transition-colors"
          aria-label="Dismiss banner"
        >
          <X className="w-3.5 h-3.5" />
        </button>
      </div>
    </div>
  );
};
