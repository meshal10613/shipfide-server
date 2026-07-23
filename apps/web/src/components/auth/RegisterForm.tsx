'use client';

import { useState } from 'react';
import { useRegisterMutation } from '@/store/api/authApi';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import Image from 'next/image';
import { ArrowRight, Lock, Mail, User, Phone, AlertCircle } from 'lucide-react';
import { motion } from 'framer-motion';
import { Button } from '@/components/ui/Button';
import { Input } from '@/components/ui/Input';

export function RegisterForm() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [phone, setPhone] = useState('');
  const [password, setPassword] = useState('');
  const [errorMsg, setErrorMsg] = useState('');
  const [register, { isLoading }] = useRegisterMutation();
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setErrorMsg('');
    try {
      const res = await register({ name, email, phone, password }).unwrap();
      if (res.success) {
        router.push(`/auth/verify-email?email=${encodeURIComponent(email)}`);
      }
    } catch (err: any) {
      setErrorMsg(err?.data?.message || 'Registration failed. Please check inputs.');
    }
  };

  return (
    <div className="w-full max-w-md space-y-4">
      <motion.div
        initial={{ opacity: 0, scale: 0.95 }}
        animate={{ opacity: 1, scale: 1 }}
        className="w-full p-8 rounded-3xl bg-white border border-slate-200 shadow-xl space-y-6"
      >
        <div className="text-center space-y-3">
          <Link href="/" className="inline-block">
            <Image
              src="/logo.png"
              alt="Shipfide Logo"
              width={180}
              height={50}
              className="h-12 w-auto object-contain mx-auto"
              priority
            />
          </Link>
          <div>
            <h1 className="text-2xl font-extrabold tracking-tight text-[#404042]">Create Account</h1>
            <p className="text-xs text-slate-500 font-medium">Join Shipfide Courier Platform</p>
          </div>
        </div>

        {errorMsg && (
          <div className="p-3 rounded-2xl bg-red-50 border border-red-200 text-[#DD0033] text-xs flex items-center gap-2">
            <AlertCircle className="w-4 h-4 shrink-0" />
            <span>{errorMsg}</span>
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-4">
          <Input
            label="Full Name"
            type="text"
            required
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="John Doe"
            leftIcon={<User className="w-4 h-4" />}
          />

          <Input
            label="Email Address"
            type="email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="name@example.com"
            leftIcon={<Mail className="w-4 h-4" />}
          />

          <Input
            label="Phone Number (Optional)"
            type="tel"
            value={phone}
            onChange={(e) => setPhone(e.target.value)}
            placeholder="+8801700000000"
            leftIcon={<Phone className="w-4 h-4" />}
          />

          <Input
            label="Password"
            type="password"
            required
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Password@123"
            leftIcon={<Lock className="w-4 h-4" />}
          />

          <Button
            type="submit"
            isLoading={isLoading}
            rightIcon={<ArrowRight className="w-4 h-4" />}
            className="w-full py-3.5 mt-2"
          >
            Register
          </Button>
        </form>

        <div className="text-center pt-2 border-t border-slate-100">
          <span className="text-xs text-slate-500">Already registered? </span>
          <Link href="/auth/login" className="text-xs font-bold text-[#DD0033] hover:underline">
            Sign In
          </Link>
        </div>
      </motion.div>

      {/* Manual Bottom Text */}
      <p className="text-center text-[11px] text-slate-500 font-medium">
        Shipfide • Delivering Trust, Every Time
      </p>
    </div>
  );
}
