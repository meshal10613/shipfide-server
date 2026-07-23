'use client';

import React from 'react';

export interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
  leftIcon?: React.ReactNode;
  rightIcon?: React.ReactNode;
}

export const Input = React.forwardRef<HTMLInputElement, InputProps>(
  ({ label, error, leftIcon, rightIcon, className = '', ...props }, ref) => {
    return (
      <div className="space-y-1.5 w-full">
        {label && (
          <label className="text-xs font-semibold text-[#404042] block">
            {label}
          </label>
        )}
        <div className="relative flex items-center">
          {leftIcon && (
            <div className="absolute left-3.5 text-slate-400 pointer-events-none">
              {leftIcon}
            </div>
          )}
          <input
            ref={ref}
            className={`w-full py-3 rounded-2xl bg-white border text-xs text-[#404042] placeholder-slate-400 transition-all focus:outline-none focus:border-[#DD0033] focus:ring-2 focus:ring-[#DD0033]/10 ${
              leftIcon ? 'pl-10' : 'pl-4'
            } ${rightIcon ? 'pr-10' : 'pr-4'} ${
              error ? 'border-red-500 focus:border-red-500' : 'border-slate-200'
            } ${className}`}
            {...props}
          />
          {rightIcon && (
            <div className="absolute right-3.5 text-slate-400">
              {rightIcon}
            </div>
          )}
        </div>
        {error && <p className="text-[11px] font-medium text-red-500 mt-1">{error}</p>}
      </div>
    );
  }
);

Input.displayName = 'Input';
