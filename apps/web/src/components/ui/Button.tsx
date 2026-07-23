'use client';

import React from 'react';
import { Loader2 } from 'lucide-react';

export interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary' | 'outline' | 'ghost' | 'danger';
  size?: 'sm' | 'md' | 'lg';
  isLoading?: boolean;
  leftIcon?: React.ReactNode;
  rightIcon?: React.ReactNode;
}

export const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  (
    {
      children,
      variant = 'primary',
      size = 'md',
      isLoading = false,
      leftIcon,
      rightIcon,
      className = '',
      disabled,
      ...props
    },
    ref
  ) => {
    const baseStyles =
      'inline-flex items-center justify-center font-bold tracking-tight rounded-full h-10 transition-all duration-200 focus:outline-none disabled:opacity-50 disabled:cursor-not-allowed select-none cursor-pointer';

    const variants = {
      primary: 'bg-[#DD0033] hover:bg-[#B30028] text-white shadow-md shadow-[#DD0033]/20 hover:scale-[1.02] active:scale-[0.98]',
      secondary: 'bg-[#404042] hover:bg-[#2B2B2D] text-white border border-[#404042]/10',
      outline: 'bg-white border border-[#E5E7EB] hover:border-[#DD0033]/40 text-[#404042] hover:text-[#DD0033]',
      ghost: 'bg-transparent hover:bg-slate-100 text-[#404042] hover:text-[#DD0033]',
      danger: 'bg-red-50 hover:bg-red-100 text-[#DD0033] border border-red-200',
    };

    const sizes = {
      sm: 'px-3.5 text-xs gap-1.5',
      md: 'px-5 text-xs gap-2',
      lg: 'px-7 text-sm gap-2.5',
    };

    return (
      <button
        ref={ref}
        disabled={disabled || isLoading}
        className={`${baseStyles} ${variants[variant]} ${sizes[size]} ${className}`}
        {...props}
      >
        {isLoading ? (
          <Loader2 className="w-4 h-4 animate-spin text-current" />
        ) : (
          leftIcon
        )}
        <span>{children}</span>
        {!isLoading && rightIcon}
      </button>
    );
  }
);

Button.displayName = 'Button';
