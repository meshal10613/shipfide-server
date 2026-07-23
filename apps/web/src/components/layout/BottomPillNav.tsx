'use client';

import { Home, ClipboardList, Bell, User } from 'lucide-react';
import Link from 'next/link';
import { usePathname } from 'next/navigation';

export function BottomPillNav() {
  const pathname = usePathname();

  const navItems = [
    { name: 'Home', href: '/dashboard', icon: Home },
    { name: 'Orders', href: '/dashboard/orders', icon: ClipboardList },
    { name: 'Alerts', href: '/dashboard/notifications', icon: Bell },
    { name: 'Profile', href: '/dashboard/profile', icon: User },
  ];

  return (
    <div className="fixed bottom-6 left-1/2 -translate-x-1/2 z-50">
      <div className="flex items-center gap-1.5 p-2 rounded-full bg-white/95 backdrop-blur-2xl border border-slate-200 shadow-xl">
        {navItems.map((item) => {
          const isActive = pathname === item.href || (item.href !== '/dashboard' && pathname.startsWith(item.href));
          return (
            <Link
              key={item.name}
              href={item.href}
              className={`flex items-center gap-2 px-4 py-2.5 rounded-full text-xs font-semibold transition-all duration-300 ${
                isActive
                  ? 'bg-[#DD0033] text-white font-extrabold shadow-md shadow-[#DD0033]/30 scale-105'
                  : 'text-slate-500 hover:text-[#404042] hover:bg-slate-100'
              }`}
            >
              <item.icon className="w-4 h-4" />
              {isActive && <span>{item.name}</span>}
            </Link>
          );
        })}
      </div>
    </div>
  );
}
