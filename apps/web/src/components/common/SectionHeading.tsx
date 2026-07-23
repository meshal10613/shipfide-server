import React from 'react';

interface SectionHeadingProps {
  badge?: string;
  title: string;
  description?: string;
  align?: 'left' | 'center' | 'right';
  className?: string;
}

export const SectionHeading: React.FC<SectionHeadingProps> = ({
  badge,
  title,
  description,
  align = 'center',
  className = '',
}) => {
  const alignClasses = {
    left: 'text-left items-start',
    center: 'text-center items-center mx-auto',
    right: 'text-right items-end ml-auto',
  };

  return (
    <div className={`flex flex-col max-w-3xl space-y-3 ${alignClasses[align]} ${className}`}>
      {badge && (
        <span className="inline-flex items-center px-3.5 py-1 rounded-full text-[11px] font-bold tracking-widest uppercase bg-[#DD0033]/10 text-[#DD0033] border border-[#DD0033]/20 shadow-xs">
          {badge}
        </span>
      )}
      <h2 className="text-3xl sm:text-4xl md:text-5xl font-black text-[#404042] tracking-tight leading-[1.15]">
        {title}
      </h2>
      {description && (
        <p className="text-base sm:text-lg text-slate-600 font-normal leading-relaxed">
          {description}
        </p>
      )}
    </div>
  );
};
