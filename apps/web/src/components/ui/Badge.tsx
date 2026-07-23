import React from 'react';

export interface BadgeProps {
  children: React.ReactNode;
  variant?: 'red' | 'blue' | 'purple' | 'amber' | 'emerald' | 'gray';
  size?: 'sm' | 'md';
  className?: string;
}

export function Badge({
  children,
  variant = 'red',
  size = 'sm',
  className = '',
}: BadgeProps) {
  const variants = {
    red: 'bg-[#DD0033]/10 text-[#DD0033] border-[#DD0033]/20',
    blue: 'bg-blue-50 text-blue-700 border-blue-200',
    purple: 'bg-purple-50 text-purple-700 border-purple-200',
    amber: 'bg-amber-50 text-amber-700 border-amber-200',
    emerald: 'bg-emerald-50 text-emerald-700 border-emerald-200',
    gray: 'bg-slate-100 text-[#404042] border-slate-200',
  };

  const sizes = {
    sm: 'px-2.5 py-0.5 text-[10px]',
    md: 'px-3 py-1 text-xs',
  };

  return (
    <span
      className={`inline-flex items-center font-mono font-bold tracking-wider rounded-full border ${variants[variant]} ${sizes[size]} ${className}`}
    >
      {children}
    </span>
  );
}
