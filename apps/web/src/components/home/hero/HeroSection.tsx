"use client";

import { motion } from "framer-motion";
import { QuickTrackingWidget } from "./QuickTrackingWidget";
import {
    ArrowRight,
    Play,
    Truck,
    CheckCircle2,
    Zap,
} from "lucide-react";
import Link from "next/link";
import { Button } from "@/components/ui/Button";

export const HeroSection = () => {
    return (
        <section className="relative overflow-hidden bg-[#F9FCFE] pt-12 pb-20 lg:pt-20 lg:pb-32 border-b border-slate-200/60">
            {/* Background Decorative Mesh Gradients */}
            <div className="absolute top-0 left-1/2 -translate-x-1/2 w-full max-w-7xl h-full pointer-events-none overflow-hidden">
                <div className="absolute -top-32 left-1/4 w-96 h-96 bg-[#DD0033]/5 rounded-full blur-3xl" />
                <div className="absolute top-1/3 right-10 w-80 h-80 bg-slate-400/5 rounded-full blur-3xl" />
            </div>

            <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative z-10">
                <div className="grid grid-cols-1 lg:grid-cols-12 gap-12 lg:gap-8 items-center">
                    {/* Left Column: Copy & Form */}
                    <motion.div
                        initial={{ opacity: 0, y: 20 }}
                        animate={{ opacity: 1, y: 0 }}
                        transition={{ duration: 0.6 }}
                        className="lg:col-span-7 space-y-8 text-center lg:text-left"
                    >
                        {/* Pill Badge */}
                        <div className="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-white border border-slate-200 shadow-xs text-xs font-bold text-[#404042]">
                            <span className="w-2 h-2 rounded-full bg-[#DD0033] animate-pulse" />
                            <span>
                                Next-Gen Logistics Engine for Bangladesh
                            </span>
                            <span className="text-slate-300">|</span>
                            <span className="text-[#DD0033]">
                                Enterprise Ready
                            </span>
                        </div>

                        {/* Headline */}
                        <h1 className="text-4xl sm:text-5xl lg:text-6xl font-black text-[#404042] tracking-tight leading-[1.1]">
                            Shipping Made{" "}
                            <span className="text-[#DD0033]">Faster.</span>{" "}
                            <br className="hidden sm:block" />
                            Business Made{" "}
                            <span className="underline decoration-[#DD0033]/30 underline-offset-8">
                                Simpler.
                            </span>
                        </h1>

                        {/* Subtext */}
                        <p className="text-base sm:text-lg text-slate-600 max-w-2xl mx-auto lg:mx-0 leading-relaxed">
                            Shipfide powers thousands of merchants with
                            nationwide last-mile parcel delivery, same-day COD
                            cashouts, smart hub warehousing, and real-time AI
                            tracking.
                        </p>

                        {/* CTA Action Buttons */}
                        <div className="flex flex-col sm:flex-row items-center justify-center lg:justify-start gap-4">
                            <Link href="/auth/register" className="w-full sm:w-auto">
                                <Button
                                    variant="primary"
                                    size="lg"
                                    rightIcon={<ArrowRight className="w-4 h-4" />}
                                    className="w-full sm:w-auto"
                                >
                                    Create Merchant Account
                                </Button>
                            </Link>
                            <Link href="/pricing" className="w-full sm:w-auto">
                                <Button
                                    variant="outline"
                                    size="lg"
                                    leftIcon={
                                        <Play className="w-4 h-4 text-[#DD0033] fill-[#DD0033]" />
                                    }
                                    className="w-full sm:w-auto"
                                >
                                    Calculate Rates
                                </Button>
                            </Link>
                        </div>

                        {/* Quick Tracking Widget Box */}
                        <div className="pt-4 flex justify-center lg:justify-start">
                            <QuickTrackingWidget />
                        </div>
                    </motion.div>

                    {/* Right Column: Hero Interactive Graphic Card */}
                    <motion.div
                        initial={{ opacity: 0, scale: 0.95 }}
                        animate={{ opacity: 1, scale: 1 }}
                        transition={{ duration: 0.7, delay: 0.2 }}
                        className="lg:col-span-5 relative"
                    >
                        {/* Outer Decorative Card */}
                        <div className="relative mx-auto max-w-md lg:max-w-none bg-white rounded-3xl border border-slate-200 shadow-2xl p-6 space-y-6">
                            {/* Top Header Row */}
                            <div className="flex items-center justify-between pb-4 border-b border-slate-100">
                                <div className="flex items-center gap-3">
                                    <div className="w-10 h-10 rounded-2xl bg-[#DD0033]/10 text-[#DD0033] flex items-center justify-center font-bold">
                                        <Truck className="w-5 h-5" />
                                    </div>
                                    <div>
                                        <h3 className="text-xs font-bold text-[#404042]">
                                            Live Delivery Dispatch
                                        </h3>
                                        <p className="text-[10px] text-slate-400 font-mono">
                                            SF-20260722-04829
                                        </p>
                                    </div>
                                </div>
                                <span className="px-3 py-1 rounded-full bg-emerald-50 text-emerald-600 text-[10px] font-bold border border-emerald-200">
                                    IN TRANSIT
                                </span>
                            </div>

                            {/* Progress Stepper */}
                            <div className="space-y-4">
                                <div className="flex items-center justify-between text-xs font-semibold text-[#404042]">
                                    <span>Hub Sorting: Central Hub</span>
                                    <span className="text-[#DD0033]">
                                        94% Completed
                                    </span>
                                </div>
                                <div className="w-full h-2.5 rounded-full bg-slate-100 overflow-hidden">
                                    <div className="h-full bg-gradient-to-r from-[#DD0033] to-rose-400 rounded-full w-[94%]" />
                                </div>
                            </div>

                            {/* Mini Stats Card */}
                            <div className="grid grid-cols-2 gap-3 pt-2">
                                <div className="p-3.5 rounded-2xl bg-slate-50 border border-slate-100">
                                    <span className="text-[10px] font-semibold text-slate-400 block">
                                        Est. Delivery
                                    </span>
                                    <span className="text-sm font-black text-[#404042]">
                                        Today, 4:30 PM
                                    </span>
                                </div>
                                <div className="p-3.5 rounded-2xl bg-slate-50 border border-slate-100">
                                    <span className="text-[10px] font-semibold text-slate-400 block">
                                        COD Collectible
                                    </span>
                                    <span className="text-sm font-black text-[#DD0033]">
                                        ৳ 2,450.00
                                    </span>
                                </div>
                            </div>

                            {/* Floating Badge 1 */}
                            <motion.div
                                animate={{ y: [0, -6, 0] }}
                                transition={{
                                    duration: 4,
                                    repeat: Infinity,
                                    ease: "easeInOut",
                                }}
                                className="absolute -top-6 -left-6 bg-white p-3.5 rounded-2xl border border-slate-200 shadow-xl flex items-center gap-3 hidden sm:flex"
                            >
                                <div className="w-8 h-8 rounded-xl bg-emerald-100 text-emerald-600 flex items-center justify-center">
                                    <CheckCircle2 className="w-4 h-4" />
                                </div>
                                <div>
                                    <span className="text-[10px] font-bold text-slate-400 block">
                                        COD Disbursement
                                    </span>
                                    <span className="text-xs font-black text-[#404042]">
                                        Same-Day Auto Payout
                                    </span>
                                </div>
                            </motion.div>

                            {/* Floating Badge 2 */}
                            <motion.div
                                animate={{ y: [0, 6, 0] }}
                                transition={{
                                    duration: 4.5,
                                    repeat: Infinity,
                                    ease: "easeInOut",
                                    delay: 0.5,
                                }}
                                className="absolute -bottom-6 -right-6 bg-[#404042] text-white p-3.5 rounded-2xl border border-slate-700 shadow-xl flex items-center gap-3 hidden sm:flex"
                            >
                                <Zap className="w-5 h-5 text-[#DD0033]" />
                                <div>
                                    <span className="text-[10px] font-semibold text-slate-400 block">
                                        Fraud Shield
                                    </span>
                                    <span className="text-xs font-bold text-white">
                                        0% Refusal Rate
                                    </span>
                                </div>
                            </motion.div>
                        </div>
                    </motion.div>
                </div>
            </div>
        </section>
    );
};
