import React from "react";
import { Navbar } from "@/components/layout/Navbar";
import { FooterSection } from "@/components/home/footer/FooterSection";

export default function CommonLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <div className="min-h-screen bg-[#F9FCFE] text-[#404042] flex flex-col justify-between selection:bg-[#DD0033] selection:text-white">
            <Navbar />
            <div>
                <main className="flex-1">{children}</main>
            </div>
            <FooterSection />
        </div>
    );
}
