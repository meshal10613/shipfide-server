"use client";

import Link from "next/link";
import Image from "next/image";
import { usePathname } from "next/navigation";
import {
    User as UserIcon,
    Menu,
    X,
    Home,
    Search,
    Calculator,
    Package,
    ShieldAlert,
    ChevronRight,
    Sparkles,
} from "lucide-react";
import { useState, useEffect } from "react";
import { createPortal } from "react-dom";
import { motion, AnimatePresence } from "framer-motion";
import { Button } from "@/components/ui/Button";

interface NavbarProps {
    currentRole?: string;
    onRoleChange?: (role: string) => void;
    userEmail?: string;
}

export function Navbar({
    currentRole = "MERCHANT",
    onRoleChange,
    userEmail,
}: NavbarProps) {
    const pathname = usePathname();
    const [showRoleDropdown, setShowRoleDropdown] = useState(false);
    const [mobileMenuOpen, setMobileMenuOpen] = useState(false);
    const [mounted, setMounted] = useState(false);

    useEffect(() => {
        setMounted(true);
    }, []);

    const centerNavLinks = [
        { href: "/", label: "Home", icon: Home },
        { href: "/tracking", label: "Tracking", icon: Search },
        { href: "/pricing", label: "Pricing", icon: Calculator },
        { href: "/guest-delivery", label: "Guest Delivery", icon: Package },
        { href: "/receiver-fraud-check", label: "Fraud Check", icon: ShieldAlert },
    ];

    const roles = [
        {
            id: "SUPER_ADMIN",
            name: "Super Admin",
            badge: "bg-purple-50 text-purple-700 border-purple-200",
        },
        {
            id: "ADMIN",
            name: "Hub Staff / Admin",
            badge: "bg-blue-50 text-blue-700 border-blue-200",
        },
        {
            id: "MERCHANT",
            name: "Merchant",
            badge: "bg-emerald-50 text-emerald-700 border-emerald-200",
        },
        {
            id: "RIDER",
            name: "Delivery Rider",
            badge: "bg-amber-50 text-amber-700 border-amber-200",
        },
    ];

    return (
        <>
            <header className="sticky top-0 z-40 bg-white/90 backdrop-blur-xl border-b border-slate-200 px-4 sm:px-8 py-3">
                <div className="max-w-7xl mx-auto flex items-center justify-between">
                    {/* Left: Brand Logo */}
                    <Link href="/" className="flex items-center group shrink-0">
                        <Image
                            src="/logo.png"
                            alt="Shipfide Logo"
                            width={160}
                            height={44}
                            className="h-10 w-auto object-contain transition-transform duration-200 group-hover:scale-105"
                            priority
                        />
                    </Link>

                    {/* Center: Desktop Navigation Routes */}
                    <nav className="hidden lg:flex items-center gap-1 bg-slate-50/80 p-1.5 rounded-full border border-slate-200/80">
                        {centerNavLinks.map((link) => {
                            const isActive = pathname === link.href;
                            return (
                                <Link key={link.href} href={link.href}>
                                    <Button
                                        variant={isActive ? "primary" : "ghost"}
                                        size="sm"
                                        className={
                                            isActive
                                                ? "shadow-sm"
                                                : "text-[#404042] hover:text-[#DD0033]"
                                        }
                                    >
                                        {link.label}
                                    </Button>
                                </Link>
                            );
                        })}
                    </nav>

                    {/* Right: Role Switcher & Auth Actions */}
                    <div className="hidden sm:flex items-center gap-3">
                        {/* Role Switcher Pill */}
                        {onRoleChange && (
                            <div className="relative">
                                <Button
                                    variant="outline"
                                    size="sm"
                                    onClick={() =>
                                        setShowRoleDropdown(!showRoleDropdown)
                                    }
                                    leftIcon={
                                        <span className="w-2 h-2 rounded-full bg-[#DD0033] animate-pulse" />
                                    }
                                >
                                    <span className="text-slate-500 font-normal">View as:</span>
                                    <span className="text-[#DD0033] font-bold ml-1">
                                        {
                                            roles.find((r) => r.id === currentRole)
                                                ?.name
                                        }
                                    </span>
                                </Button>

                                {showRoleDropdown && (
                                    <div className="absolute right-0 mt-2 w-56 rounded-2xl bg-white border border-slate-200 shadow-xl p-2 z-50 animate-in fade-in zoom-in-95">
                                        <div className="text-[10px] font-mono uppercase tracking-wider text-slate-400 px-3 py-1.5">
                                            Switch Dashboard View
                                        </div>
                                        {roles.map((r) => (
                                            <button
                                                key={r.id}
                                                onClick={() => {
                                                    onRoleChange(r.id);
                                                    setShowRoleDropdown(false);
                                                }}
                                                className={`w-full flex items-center justify-between px-3 py-2 rounded-xl text-xs font-medium transition-colors ${
                                                    currentRole === r.id
                                                        ? "bg-[#DD0033] text-white font-bold"
                                                        : "text-[#404042] hover:bg-slate-50"
                                                }`}
                                            >
                                                <span>{r.name}</span>
                                                <span
                                                    className={`px-2 py-0.5 rounded-full text-[10px] font-mono border ${r.badge}`}
                                                >
                                                    {r.id}
                                                </span>
                                            </button>
                                        ))}
                                    </div>
                                )}
                            </div>
                        )}

                        {/* Desktop Auth Button */}
                        <Link href="/auth/login">
                            <Button
                                variant="primary"
                                size="sm"
                                leftIcon={<UserIcon className="w-3.5 h-3.5" />}
                            >
                                Sign In
                            </Button>
                        </Link>
                    </div>

                    {/* Mobile Hamburger Toggle Button */}
                    <div className="flex items-center lg:hidden gap-2">
                        <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => setMobileMenuOpen(true)}
                            aria-label="Open mobile navigation sidebar"
                            className="w-10 h-10 p-0 rounded-full border border-slate-200 bg-slate-50"
                        >
                            <Menu className="w-5 h-5 text-[#404042]" />
                        </Button>
                    </div>
                </div>
            </header>

            {/* Portaled Mobile Slide-Over Sidebar Drawer (Appended directly to document.body) */}
            {mounted &&
                createPortal(
                    <AnimatePresence>
                        {mobileMenuOpen && (
                            <>
                                {/* Backdrop Overlay */}
                                <motion.div
                                    initial={{ opacity: 0 }}
                                    animate={{ opacity: 1 }}
                                    exit={{ opacity: 0 }}
                                    onClick={() => setMobileMenuOpen(false)}
                                    className="fixed inset-0 bg-slate-950/60 backdrop-blur-xs z-[9999] lg:hidden"
                                />

                                {/* Slide-over Panel */}
                                <motion.aside
                                    initial={{ x: "100%" }}
                                    animate={{ x: 0 }}
                                    exit={{ x: "100%" }}
                                    transition={{
                                        type: "spring",
                                        damping: 25,
                                        stiffness: 200,
                                    }}
                                    className="fixed top-0 right-0 bottom-0 w-80 max-w-[85vw] bg-white z-[9999] shadow-2xl p-6 flex flex-col justify-between overflow-y-auto lg:hidden"
                                >
                                    <div className="space-y-6">
                                        {/* Sidebar Top Header */}
                                        <div className="flex items-center justify-between pb-4 border-b border-slate-100">
                                            <Link
                                                href="/"
                                                onClick={() => setMobileMenuOpen(false)}
                                                className="inline-block"
                                            >
                                                <Image
                                                    src="/logo.png"
                                                    alt="Shipfide Logo"
                                                    width={140}
                                                    height={40}
                                                    className="h-9 w-auto object-contain"
                                                />
                                            </Link>
                                            <Button
                                                variant="ghost"
                                                size="sm"
                                                onClick={() => setMobileMenuOpen(false)}
                                                aria-label="Close sidebar"
                                                className="w-9 h-9 p-0 rounded-full border border-slate-200"
                                            >
                                                <X className="w-4 h-4 text-[#404042]" />
                                            </Button>
                                        </div>

                                        {/* Prominent Sign In Section in Mobile Sidebar */}
                                        <div className="bg-gradient-to-r from-rose-50 to-slate-50 p-4 rounded-3xl border border-rose-100 space-y-3">
                                            <div className="flex items-center gap-2">
                                                <Sparkles className="w-4 h-4 text-[#DD0033]" />
                                                <span className="text-xs font-bold text-[#404042]">
                                                    Account Access
                                                </span>
                                            </div>
                                            <Link
                                                href="/auth/login"
                                                onClick={() => setMobileMenuOpen(false)}
                                                className="block w-full"
                                            >
                                                <Button
                                                    variant="primary"
                                                    size="md"
                                                    leftIcon={<UserIcon className="w-4 h-4" />}
                                                    className="w-full shadow-md shadow-[#DD0033]/20"
                                                >
                                                    Sign In
                                                </Button>
                                            </Link>
                                            <div className="text-center pt-1">
                                                <Link
                                                    href="/auth/register"
                                                    onClick={() => setMobileMenuOpen(false)}
                                                    className="text-xs font-bold text-[#DD0033] hover:underline"
                                                >
                                                    Don't have an account? Register
                                                </Link>
                                            </div>
                                        </div>

                                        {/* Mobile Center Routes List */}
                                        <div className="space-y-1">
                                            <div className="text-[10px] font-mono uppercase tracking-wider text-slate-400 px-3 py-1">
                                                Navigation Menu
                                            </div>
                                            {centerNavLinks.map((link) => {
                                                const isActive = pathname === link.href;
                                                const Icon = link.icon;
                                                return (
                                                    <Link
                                                        key={link.href}
                                                        href={link.href}
                                                        onClick={() => setMobileMenuOpen(false)}
                                                        className="block"
                                                    >
                                                        <div
                                                            className={`flex items-center justify-between px-3.5 py-3 rounded-2xl text-xs font-bold transition-all ${
                                                                isActive
                                                                    ? "bg-[#DD0033] text-white shadow-md shadow-[#DD0033]/20"
                                                                    : "text-[#404042] hover:bg-slate-50"
                                                            }`}
                                                        >
                                                            <div className="flex items-center gap-3">
                                                                <Icon className="w-4 h-4 shrink-0" />
                                                                <span>{link.label}</span>
                                                            </div>
                                                            <ChevronRight className={`w-3.5 h-3.5 opacity-60 ${isActive ? "text-white" : "text-slate-400"}`} />
                                                        </div>
                                                    </Link>
                                                );
                                            })}
                                        </div>

                                        {/* Role Switcher in Sidebar */}
                                        {onRoleChange && (
                                            <div className="space-y-2 pt-2 border-t border-slate-100">
                                                <div className="text-[10px] font-mono uppercase tracking-wider text-slate-400 px-3 py-1">
                                                    Switch Role View
                                                </div>
                                                <div className="grid grid-cols-1 gap-1.5">
                                                    {roles.map((r) => (
                                                        <Button
                                                            key={r.id}
                                                            variant={
                                                                currentRole === r.id
                                                                    ? "primary"
                                                                    : "outline"
                                                            }
                                                            size="sm"
                                                            onClick={() => {
                                                                onRoleChange(r.id);
                                                                setMobileMenuOpen(false);
                                                            }}
                                                            className="w-full justify-between h-9 text-xs"
                                                        >
                                                            <span>{r.name}</span>
                                                            <span className="text-[10px] opacity-75 font-mono">
                                                                {r.id}
                                                            </span>
                                                        </Button>
                                                    ))}
                                                </div>
                                            </div>
                                        )}
                                    </div>

                                    {/* Sidebar Footer */}
                                    <div className="pt-6 border-t border-slate-100 text-center space-y-1">
                                        <p className="text-xs font-bold text-[#404042]">
                                            Shipfide Logistics Platform
                                        </p>
                                        <p className="text-[10px] text-slate-400">
                                            Delivering Trust, Every Time • © 2026
                                        </p>
                                    </div>
                                </motion.aside>
                            </>
                        )}
                    </AnimatePresence>,
                    document.body
                )}
        </>
    );
}
