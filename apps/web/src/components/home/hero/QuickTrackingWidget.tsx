"use client";

import { useState } from "react";
import { Search, ArrowRight, ShieldCheck } from "lucide-react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/Button";

export const QuickTrackingWidget = () => {
    const [trackingId, setTrackingId] = useState("");
    const router = useRouter();

    const handleTrack = (e: React.FormEvent) => {
        e.preventDefault();
        if (!trackingId.trim()) return;
        router.push(`/tracking?code=${encodeURIComponent(trackingId.trim())}`);
    };

    return (
        <div className="bg-white p-3 sm:p-4 rounded-3xl border border-slate-200 shadow-xl shadow-slate-200/60 max-w-xl w-full">
            <form
                onSubmit={handleTrack}
                className="flex flex-col sm:flex-row items-center gap-2"
            >
                <div className="relative w-full flex-1">
                    <Search className="w-5 h-5 text-slate-400 absolute left-4 top-1/2 -translate-y-1/2" />
                    <input
                        type="text"
                        value={trackingId}
                        onChange={(e) => setTrackingId(e.target.value)}
                        placeholder="Enter Tracking ID (e.g. SF-20260722-04829)"
                        className="w-full pl-11 pr-4 py-2.5 rounded-full bg-slate-50 text-xs sm:text-sm font-mono font-medium text-[#404042] placeholder:text-slate-400 border border-slate-200 focus:outline-none focus:border-[#DD0033] focus:bg-white transition-all"
                    />
                </div>
                <Button
                    type="submit"
                    variant="primary"
                    rightIcon={<ArrowRight className="w-4 h-4" />}
                    className="w-full sm:w-auto"
                >
                    Track Order
                </Button>
            </form>
            <div className="flex items-center justify-between mt-3 px-2 text-[11px] text-slate-400 font-medium">
                <span className="flex items-center gap-1">
                    <ShieldCheck className="w-3.5 h-3.5 text-emerald-600" />
                    Instant Live Telemetry
                </span>
                <span>64 Districts Covered</span>
            </div>
        </div>
    );
};
