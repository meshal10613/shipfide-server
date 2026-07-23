export interface FeatureItem {
  id: string;
  title: string;
  description: string;
  iconName: string;
  badge?: string;
}

export interface ServiceItem {
  id: string;
  title: string;
  subtitle: string;
  description: string;
  features: string[];
  metrics: string;
  imageAlt: string;
}

export interface PricingPlan {
  id: string;
  name: string;
  description: string;
  priceMonthly: number;
  priceYearly: number;
  isPopular?: boolean;
  features: string[];
  ctaText: string;
}

export interface TestimonialItem {
  id: string;
  name: string;
  role: string;
  company: string;
  avatar: string;
  quote: string;
  rating: number;
  shipmentCount: string;
}

export interface FaqItem {
  id: string;
  question: string;
  answer: string;
  category: string;
}

export const ANNOUNCEMENT_DATA = {
  badge: 'NEW RELEASE',
  title: 'Shipfide API v2.0 & Automated Linehaul Routing is live!',
  ctaText: 'Explore Developer Docs →',
  ctaUrl: '/dashboard/support',
};

export const FEATURES_DATA: FeatureItem[] = [
  {
    id: 'tracking',
    title: 'Real-Time GPS Tracking',
    description: 'Sub-second parcel telemetry across 64 districts with instant SMS & WhatsApp notifications for end customers.',
    iconName: 'Navigation',
    badge: 'AI Telemetry',
  },
  {
    id: 'fleet',
    title: 'Fleet & Driver Roster Management',
    description: 'Automated shift scheduling, vehicle maintenance telemetry, and real-time rider capacity allocation.',
    iconName: 'Truck',
    badge: 'Operations',
  },
  {
    id: 'warehouse',
    title: 'Smart Warehouse Fulfillment',
    description: 'Zone-based inventory sorting, barcode scanning at hub inbound, and automated parcel grouping.',
    iconName: 'Building2',
  },
  {
    id: 'dashboard',
    title: 'Executive Analytics Dashboard',
    description: 'Full visibility over COD collections, delivery success rates, return ratios, and merchant payouts.',
    iconName: 'BarChart3',
    badge: 'Live Data',
  },
  {
    id: 'fraud-check',
    title: 'Receiver Fraud & Risk Engine',
    description: 'Cross-platform AI risk detection to identify chronic returners and non-delivering addresses before dispatch.',
    iconName: 'ShieldAlert',
    badge: 'Exclusive',
  },
  {
    id: 'cod-settlement',
    title: 'Same-Day COD Cashouts',
    description: 'Instant automated cash disbursement to bKash, Nagad, or BEFTN Bank accounts upon delivery confirmation.',
    iconName: 'Wallet',
  },
  {
    id: 'api-integration',
    title: 'Developer REST & GraphQL APIs',
    description: 'Plug-and-play SDKs for Shopify, WooCommerce, custom Next.js stores, and ERP systems.',
    iconName: 'Code2',
  },
  {
    id: 'bulk-booking',
    title: 'Bulk CSV & Excel Booking',
    description: 'Import 5,000+ orders in seconds with auto-validation for postal codes and phone numbers.',
    iconName: 'FileSpreadsheet',
  },
  {
    id: 'route-optimization',
    title: 'AI Route Optimization',
    description: 'Reduce last-mile delivery costs by 28% through dynamic traffic-aware rider dispatching.',
    iconName: 'Zap',
    badge: '28% Faster',
  },
];

export const SERVICES_DATA: ServiceItem[] = [
  {
    id: 'express',
    title: 'Express Doorstep Delivery',
    subtitle: 'Same-Day & Next-Day Last Mile Delivery',
    description: 'Guaranteed 24-hour delivery inside Dhaka and 48-hour nationwide delivery with OTP customer handover.',
    features: ['Instant OTP Verification', 'Live Driver GPS Tracking', '3 Free Delivery Attempts', 'Doorstep COD Collection'],
    metrics: '99.8% On-Time SLA',
    imageAlt: 'Express Courier',
  },
  {
    id: 'freight',
    title: 'Inter-District Linehaul Freight',
    subtitle: 'Heavy Cargo & B2B Shipping',
    description: 'Scheduled container trucks and linehaul vans connecting regional hubs across 8 divisions daily.',
    features: ['Full Truck Load (FTL)', 'Less Than Truckload (LTL)', 'Dedicated Cargo Escort', 'Climate-Controlled Vans'],
    metrics: '64 Districts Covered',
    imageAlt: 'Freight Trucking',
  },
  {
    id: 'warehousing',
    title: 'Smart Hub Warehousing',
    subtitle: 'Storage & Pick-and-Pack Fulfillment',
    description: 'Strategic hub storage with barcode inventory tracking, automated sorting, and multi-location pickup.',
    features: ['24/7 Security CCTV', 'Climate Controlled Storage', 'Automated Barcode Sorting', 'Same-Day Pick & Pack'],
    metrics: '500,000 Sq Ft Capacity',
    imageAlt: 'Warehouse Storage',
  },
  {
    id: 'crossborder',
    title: 'Cross-Border & E-Commerce Freight',
    subtitle: 'Global Air Freight & Customs Handling',
    description: 'Streamlined international export and import shipping with automated customs clearance documentation.',
    features: ['Customs Clearance', 'Air Freight Forwarding', 'Door-to-Door Duty Handling', 'Global Track & Trace'],
    metrics: '120+ International Destinations',
    imageAlt: 'Cross Border Shipping',
  },
];

export const WHY_CHOOSE_DATA = {
  badge: 'ENTERPRISE ADVANTAGE',
  title: 'Built for High-Volume Merchants & Logistics Leaders',
  description: 'Shipfide replaces fragmented manual courier processes with a single unified operating system for logistics, fleet dispatch, and COD treasury.',
  stats: [
    { label: 'Nationwide Delivery Coverage', value: '64 Districts' },
    { label: 'Average Last-Mile Delivery Time', value: '18.4 Hours' },
    { label: 'Successful Delivery Rate', value: '99.4%' },
    { label: 'Automated Daily COD Payouts', value: '৳ 50M+' },
  ],
  points: [
    { title: 'Zero Loss Guaranteed', desc: 'Full transit insurance coverage on high-value electronics and merchant cargo.' },
    { title: 'Dedicated Key Account Managers', desc: 'Direct phone & Slack support for merchants shipping over 500 parcels/month.' },
    { title: 'Transparent Dynamic Pricing', desc: 'No hidden fuel surcharges or surprise return fees.' },
  ],
};

export const HOW_IT_WORKS_STEPS = [
  {
    step: '01',
    title: 'Book Parcel or Bulk Import',
    description: 'Create orders manually in under 10 seconds or upload thousands via CSV / API integration.',
    iconName: 'PackagePlus',
  },
  {
    step: '02',
    title: 'Rider Pickup from Warehouse',
    description: 'Automated rider assignment to collect packages directly from your specified pickup hub.',
    iconName: 'Truck',
  },
  {
    step: '03',
    title: 'Hub Sorting & Linehaul Transit',
    description: 'Barcode scanning and automated sorting into divisional transit vehicles for fast hub routing.',
    iconName: 'Building2',
  },
  {
    step: '04',
    title: 'OTP Handover & Same-Day Cashout',
    description: 'Secured delivery handover via 6-digit OTP and instant COD balance disbursement to your wallet.',
    iconName: 'CheckCircle2',
  },
];

export const TESTIMONIALS_DATA: TestimonialItem[] = [
  {
    id: 't1',
    name: 'Tanvir Ahmed',
    role: 'Head of Operations',
    company: 'TechGear BD',
    avatar: 'https://images.unsplash.com/photo-1534528741775-53994a69daeb?w=150&auto=format&fit=crop&q=80',
    quote: 'Switching to Shipfide reduced our customer delivery complaints by 75%. Same-day COD cashout is a game-changer for our cash flow.',
    rating: 5,
    shipmentCount: '12,500+ Parcels / Mo',
  },
  {
    id: 't2',
    name: 'Nusrat Jahan',
    role: 'Founder & CEO',
    company: 'Glamour Closet Bangladesh',
    avatar: 'https://images.unsplash.com/photo-1580489944761-15a19d654956?w=150&auto=format&fit=crop&q=80',
    quote: 'The receiver fraud check tool alone saved us over ৳200,000 in fake order returns last month. Unmatched reliability!',
    rating: 5,
    shipmentCount: '8,200+ Parcels / Mo',
  },
  {
    id: 't3',
    name: 'Sabbir Hossain',
    role: 'Logistics Director',
    company: 'Apex Electronics',
    avatar: 'https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=150&auto=format&fit=crop&q=80',
    quote: 'Their API integration took less than 2 hours to connect with our custom Next.js storefront. Instant tracking for every customer.',
    rating: 5,
    shipmentCount: '25,000+ Parcels / Mo',
  },
];

export const STATS_DATA = [
  { label: 'Parcels Delivered Nationwide', value: 5000000, suffix: '+', prefix: '' },
  { label: 'Active Enterprise Merchants', value: 12500, suffix: '+', prefix: '' },
  { label: 'On-Time Delivery SLA', value: 99.8, suffix: '%', prefix: '' },
  { label: 'Operational Hubs', value: 64, suffix: ' Districts', prefix: '' },
];

export const PRICING_PLANS: PricingPlan[] = [
  {
    id: 'starter',
    name: 'Starter Merchant',
    description: 'Perfect for small e-commerce stores and social media sellers.',
    priceMonthly: 60,
    priceYearly: 50,
    features: [
      'Up to 200 Parcels / month',
      'Inside Dhaka: ৳60 Base Rate',
      'Standard Next-Day Delivery',
      'Basic SMS Telemetry',
      'Weekly COD Cashout',
    ],
    ctaText: 'Start Shipping',
  },
  {
    id: 'growth',
    name: 'Growth Enterprise',
    description: 'Ideal for fast-scaling online brands requiring priority dispatch.',
    priceMonthly: 55,
    priceYearly: 45,
    isPopular: true,
    features: [
      'Up to 5,000 Parcels / month',
      'Inside Dhaka: ৳55 Base Rate',
      'Same-Day Priority Express',
      'Receiver Fraud Check Engine',
      'Daily Automated COD Cashout',
      'Dedicated Account Manager',
      'Full REST API & Webhooks Access',
    ],
    ctaText: 'Get Started with Growth',
  },
  {
    id: 'enterprise',
    name: 'Corporate Fleet',
    description: 'Custom warehousing, dedicated trucks, & volume discounts.',
    priceMonthly: 45,
    priceYearly: 38,
    features: [
      'Unlimited Parcels / month',
      'Custom Volume Discounting',
      'Dedicated Warehouse Storage',
      'FTL & LTL Linehaul Fleet',
      'Instant COD Wallet Disbursement',
      'Custom SLA & Insurance Coverage',
      '24/7 Priority Support Hotline',
    ],
    ctaText: 'Contact Enterprise Sales',
  },
];

export const FAQ_DATA: FaqItem[] = [
  {
    id: 'faq-1',
    question: 'How fast does Shipfide process COD cash disbursement to merchants?',
    answer: 'Shipfide provides automated same-day and daily COD payouts directly to your bKash, Nagad, or Bank Account as soon as the delivery rider confirms OTP handover.',
    category: 'Payments & COD',
  },
  {
    id: 'faq-2',
    question: 'What is the Receiver Fraud Check feature and how does it protect sellers?',
    answer: 'Our proprietary Receiver Fraud Check engine scans our nationwide network history to identify phone numbers with high parcel cancellation or refusal rates before you dispatch.',
    category: 'Features & Security',
  },
  {
    id: 'faq-3',
    question: 'Which areas in Bangladesh does Shipfide cover?',
    answer: 'Shipfide covers all 64 districts and 495+ upazilas across Bangladesh with direct hub operations in all major divisional cities.',
    category: 'Coverage',
  },
  {
    id: 'faq-4',
    question: 'How do I integrate Shipfide with my online e-commerce website?',
    answer: 'We offer official pre-built plugins for Shopify and WooCommerce, as well as robust REST & GraphQL APIs with SDK support for Next.js, Node.js, and Python.',
    category: 'API & Integrations',
  },
  {
    id: 'faq-5',
    question: 'What happens if a parcel is damaged or lost during transit?',
    answer: 'All Shipfide shipments include baseline transit protection up to ৳ 10,000. Enterprise plans include 100% full-value insurance coverage.',
    category: 'Insurance & Claims',
  },
];
