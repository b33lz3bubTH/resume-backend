#!/bin/bash

BASE_URL="http://localhost:8080"
ROOT_KEY="abcd123"

echo "Seeding database..."

# Create Bootcamp 1
echo "Creating Bootcamp 1..."
curl -s -X POST "$BASE_URL/api/bootcamps" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d '{"title": "Ultimate VibeCoding Backend Mastery", "subtitle": "Build Production-Ready E-Commerce with Node.js, TypeScript & Prisma", "description": "Master backend development by building a complete, config-based e-commerce platform that anyone can deploy. Learn industry-standard practices while contributing to an open-source community project.", "long_description": "This intensive bootcamp takes you from backend fundamentals to advanced production deployment. You'\''ll build a fully-functional, config-based e-commerce platform that'\''s designed for easy customization without touching code. Perfect for college students and developers looking to master modern backend development while creating something real and deployable.", "tech_stack": ["Node.js", "TypeScript", "Prisma", "PostgreSQL", "Express", "JWT", "Docker", "Golang", "n8n"], "duration": "8 Weeks", "level": "Intermediate to G0d Level", "price": "₹2,999", "highlights": ["Teacher 2 - Me (b33lz3bubTH), Teacher 1 - Chayan Maji (RyzenPython)", "Build a complete production-ready e-commerce backend", "Master TypeScript with real-world patterns", "Learn database design with Prisma ORM", "Implement authentication & authorization", "Config-based architecture for easy deployment", "Open-source contribution opportunities", "Lifetime access to course materials", "Live coding sessions & code reviews", "Community support & mentorship"], "modules": [{"title": "Backend Fundamentals & TypeScript Setup", "description": "Set up your development environment, understand Node.js architecture, and master TypeScript basics", "duration": "Week 1", "topics": ["Node.js fundamentals", "TypeScript configuration", "Project structure", "Environment setup", "Version control with Git"]}, {"title": "Database Design & Prisma Mastery", "description": "Design scalable database schemas, master Prisma ORM, and implement migrations", "duration": "Week 2", "topics": ["Database design principles", "Prisma schema", "Relations & migrations", "Query optimization", "PostgreSQL advanced features"]}, {"title": "API Development & Express Architecture", "description": "Build RESTful APIs with Express, implement middleware, and handle errors gracefully", "duration": "Week 3", "topics": ["Express.js patterns", "Middleware architecture", "Request validation", "Error handling", "API versioning"]}, {"title": "Authentication & Security", "description": "Implement JWT-based authentication, role-based access control, and security best practices", "duration": "Week 4", "topics": ["JWT authentication", "Password hashing", "RBAC implementation", "Security headers", "Rate limiting"]}, {"title": "Config-Based Architecture", "description": "Build a flexible, config-driven system that allows deployment without code changes", "duration": "Week 5", "topics": ["Configuration management", "Environment variables", "Feature flags", "Multi-tenant architecture", "Dynamic routing"]}, {"title": "Payment Integration & Order Management", "description": "Integrate payment gateways, handle orders, and implement inventory management", "duration": "Week 6", "topics": ["Payment gateway integration", "Order processing", "Inventory management", "Transaction handling", "Webhook processing"]}, {"title": "Caching, Performance & Deployment", "description": "Optimize performance with Redis, implement caching strategies, and deploy to production", "duration": "Week 7", "topics": ["Redis caching", "Query optimization", "Docker containerization", "CI/CD pipelines", "Production best practices"]}, {"title": "Community Contribution & Open Source", "description": "Learn to contribute to open-source, code reviews, and maintaining production systems", "duration": "Week 8", "topics": ["Git workflows", "Code reviews", "Documentation", "Testing strategies", "Community building"]}], "project_features": ["Multi-vendor marketplace support", "Product catalog management", "Cart & checkout system", "Order tracking & management", "User authentication & profiles", "Admin dashboard", "Payment gateway integration", "Inventory management", "Email notifications", "Config-based customization", "API documentation", "Docker deployment ready"], "target_audience": ["College students learning backend development", "Developers transitioning to TypeScript", "Anyone wanting to build production-ready systems", "Open-source contributors", "Self-learners with basic coding knowledge", "Entrepreneurs wanting to deploy their own e-commerce"], "images": ["https://plus.unsplash.com/premium_photo-1661331911412-330f2e99cf53?q=80&w=1470&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"], "videos": ["-oUxzJEmzsg"], "github_repo": "https://github.com/b33lz3bubTH/nrix7-ecommerce-layout", "demo_url": "", "status": "active", "enrolled_count": 68, "rating": 4.1}' > /dev/null

# Create Bootcamp 2
echo "Creating Bootcamp 2..."
curl -s -X POST "$BASE_URL/api/bootcamps" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d '{"title": "Low-Level Programming Mastery: x64 ASM, ABIs & System Interfaces", "subtitle": "Write code that talks directly to the machine. Serious engineering.", "description": "Advanced, hands-on bootcamp focused on x64 assembly, calling conventions, ABIs, and how to interface with operating systems and runtimes like a pro.", "long_description": "A concentrated 5-week, online-only program for advanced learners. We will go deep into x64 assembly, explore platform ABIs (System V, Windows x64), calling conventions, stack frames, linkage, and how higher-level languages talk to the OS. You'\''ll write assembly that calls libc/syscalls, interop with C, and understand what compilers actually generate.", "tech_stack": ["x64 Assembly", "System V ABI", "Windows x64 ABI", "GAS/NASM", "C interop", "Linux syscalls"], "duration": "5 Weeks", "level": "Advanced", "price": "₹3,999", "highlights": ["Instructor: Sourav Ahmed (b33lz3bubTH)", "Deep dive into calling conventions and ABIs", "Hands-on assembly with real tooling", "Interop between ASM and C libraries", "Understand ELF/PE basics and linkage", "Practical OS interface examples (Linux focus)"], "modules": [{"title": "x64 Architecture & Toolchain Setup", "description": "Registers, memory model, assemblers, linkers, and debugging tools", "duration": "Week 1", "topics": ["x64 registers", "stack & calling basics", "GAS vs NASM", "ld & objdump", "gdb"]}, {"title": "Calling Conventions & ABIs", "description": "System V vs Windows x64, argument passing, stack frames, prologue/epilogue", "duration": "Week 2", "topics": ["System V ABI", "Windows x64 ABI", "callee vs caller saved", "varargs"]}, {"title": "Syscalls & libc Interop", "description": "Making syscalls, linking with libc, calling C from ASM and vice versa", "duration": "Week 3", "topics": ["syscall interface", "libc interop", "extern symbols", "name mangling"]}, {"title": "ELF/PE & Linking Basics", "description": "Understand object files, sections, symbols, relocations, and loaders", "duration": "Week 4", "topics": ["ELF sections", "PE overview", "linker scripts", "relocations"]}, {"title": "Capstone & Optimization", "description": "Build a small utility in ASM with C interop and measure performance", "duration": "Week 5", "topics": ["capstone project", "perf basics", "micro-optimizations"]}], "project_features": ["Write pure assembly that interfaces with the OS", "Call into libc functions safely", "Understand and implement calling conventions", "Link assembly with C for practical tasks", "Inspect compiler output to learn real patterns"], "target_audience": ["Advanced programmers with systems interest", "Security researchers and RE enthusiasts", "Performance-focused engineers", "Low-level curious developers"], "images": ["https://www.virusbulletin.com/uploads/images/figures/2013/03/shellcoding-6.jpg", "https://www.corelan.be/wp-content/uploads/2010/02/image7.png"], "videos": [], "status": "upcoming", "enrolled_count": 0}' > /dev/null

# Create Journal Entry 1
echo "Creating Journal Entry 1..."
curl -s -X POST "$BASE_URL/api/journal" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d '{"title": "Beyond Code: My Journey in Real Estate Development & Full-Stack Innovation", "body": "<div class='\''vintage-content'\''>\n      <h2>Balancing Two Worlds</h2>\n      <p>While many know me as a developer, I'\''m actively involved in my family'\''s real estate business, managing construction projects while pursuing my passion for technology. This unique blend gives me a practical perspective on building things—both in the physical and digital worlds.</p>\n      \n      <h2>Real Estate Portfolio</h2>\n      <p><strong>Completed Projects:</strong><br />\n      • Sarat Colony (Ahmed Villa Apartment), Near Kolkata Airport - G+2 Building<br />\n      • Rajarhat Chowmatha (Maa Santoshi Apartment) - G+4 Building<br />\n      <br />\n      <strong>Ongoing Development:</strong><br />\n      • Kaikhali Malirbagan - G+3 Building (Currently under construction)</p>\n      \n      <h2>Technical Arsenal</h2>\n      <p><strong>Full-Stack Mastery:</strong><br />\n      Python & MERN Stack (MongoDB, Express, React, Node.js) - Building scalable web applications from database to UI</p>\n      \n      <p><strong>Systems-Level Expertise:</strong><br />\n      Linux ricing enthusiast with deep system customization experience<br />\n      Low-level systems programming and malware development research<br />\n      Understanding of system internals beyond typical web development</p>\n      \n      <h2>Notable Projects</h2>\n      <p><strong>E-Commerce Framework:</strong><br />\n      Built a complete Shopify-like platform from scratch<br />\n      Single-handedly developed a JSON-configurable e-commerce system<br />\n      Enables rapid store deployment through configuration rather than coding</p>\n      \n      <p><strong>AI-Powered WhatsApp Commerce:</strong><br />\n      Developed a custom WhatsApp bot for premium e-commerce clients<br />\n      Features include: product browsing, ordering, tracking, and customer support<br />\n      Built without GPT dependencies - no token limits or API constraints<br />\n      Provides seamless conversational commerce experience<br />\n      One-time payment model with full customization and support</p>\n      \n      <p><strong>Machine Learning Integration:</strong><br />\n      Experience with model fine-tuning for specific business use cases<br />\n      Practical AI implementation beyond theoretical knowledge</p>\n      \n      <h2>Philosophy & Approach</h2>\n      <p>My real estate experience taught me about project management, client relations, and delivering tangible results. I apply these same principles to software development—focusing on robust, maintainable solutions that solve real business problems.</p>\n      \n      <h2>Why This Matters</h2>\n      <p>This diverse background gives me a unique advantage: I understand both the technical implementation and the business impact. Whether it'\''s managing a construction timeline or architecting a scalable software system, the fundamentals of planning, execution, and delivery remain the same.</p>\n    </div>", "summary": "A look at my dual expertise in real estate development and full-stack software engineering, highlighting completed projects, technical skills, and innovative e-commerce solutions built from scratch.", "published_on": "2025-10-01", "category": "Personal Journey", "tags": ["real-estate", "full-stack", "ecommerce", "python", "mern", "linux", "whatsapp-bot", "machine-learning"], "author": "b33lz3bubTH", "read_time": "5 min read"}' > /dev/null

# Create Journal Entry 2
echo "Creating Journal Entry 2..."
curl -s -X POST "$BASE_URL/api/journal" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d '{"title": "Hackathon: Building a Gen-AI Fact-Checker that Talks on Telegram", "body": "<div class=\"vintage-content\">\n      <h2>The Brief</h2>\n      <p>The hackathon prompt was deceptively simple: \"Combat misinformation.\" In 8 hours our team built a Gen-AI powered fact-checking assistant that ingests messages from Telegram chats, pulls evidence from open web sources, and responds with concise, sourced corrections.</p>\n  \n      <h2>Why This Project</h2>\n      <p>With misinformation spreading fast on messaging apps, we wanted something low friction — a bot that could intervene politely, provide verifiable context, and link to original sources so readers can decide for themselves.</p>\n  \n      <h2>Tech Stack & Architecture</h2>\n      <p>We used an open-source foundation model for natural language understanding and claim classification, a lightweight retrieval layer that queries vetted news + fact-check indexes, and a small FastAPI microservice to orchestrate processing. The Telegram integration was a webhook-based adapter that listens to new messages and posts non-intrusive corrections when confidence was high.</p>\n  \n      <h2>Engineering Highlights</h2>\n      <p>— Real-time message parsing with minimal latency so conversations felt natural.<br />\n         — A retrieval-augmented pipeline that fetches fresh sources and ranks them by recency and reliability.<br />\n         — A tiny prompts/chain module to force the model to (1) identify the claim, (2) score factuality, (3) propose concise wording, and (4) attach 1–2 source links.</p>\n  \n      <h2>The UX</h2>\n      <p>We prioritized clarity and tone: corrections were framed as suggestions, not judgments — e.g. \"Quick check: sources show X — here are two links if you want to read more.\" Users could tap to see the full provenance and the model'\''s reasoning summary.</p>\n  \n      <h2>The Demo</h2>\n      <p>On stage we simulated a noisy Telegram group where a dubious claim went viral. The bot quietly replied with a short, evidence-backed correction and a link to a reliable source. Judges appreciated the restraint — the bot amplified context rather than policing conversations.</p>\n  \n      <h2>Lessons & Next Steps</h2>\n      <p>We learned how crucial precision and source quality are: false positives kill trust faster than silence. Next steps include better source vetting, community opt-in, and a transparent appeals workflow so users can challenge the bot'\''s corrections.</p>\n  \n      <p>Even without winning the top prize, the project felt meaningful — a small, pragmatic step toward improving information hygiene in the apps people actually use.</p>\n    </div>", "summary": "Built a Telegram-integrated Gen-AI fact-checking bot during a 8-hour hackathon. Focused on retrieval-augmented verification, polite UX, and transparent sourcing to reduce misinformation in chat groups.", "published_on": "2025-09-22"}' > /dev/null

# Create Journal Entry 3
echo "Creating Journal Entry 3..."
curl -s -X POST "$BASE_URL/api/journal" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d '{"title": "Offensive Research: Cross-Platform Media-Bound Persistence Threat (Hackathon Project)", "body": "<div class='\''vintage-content'\''>\n      <h2>Scope & Objectives</h2>\n      <p>For an upcoming security hackathon, I developed a sophisticated cross-platform threat that leverages user trust in media files. The core concept involves a custom executable that masquerades as a video file, complete with manipulated thumbnails and metadata. When executed, it simultaneously plays legitimate media content while deploying polymorphic, fully undetectable (FUD) malware in the background.</p>\n      \n      <h2>Target & Strategy</h2>\n      <p>The objective was to create a dual-purpose executable that maintains perfect disguise while achieving persistent compromise. The program stores an actual media file internally, displays expected video playback to the user, and concurrently executes polymorphic shellcode that evades traditional detection methods. Persistence mechanisms ensure the malware survives reboots and remains active across Windows, macOS, and Linux environments.</p>\n      \n      <h2>Development Process</h2>\n      <p>1. Engineered a custom launcher that bundles media files and malicious payloads within a single executable<br />\n         2. Implemented thumbnail spoofing and metadata manipulation to enhance social engineering<br />\n         3. Developed polymorphic code generation to create signature-evading payloads<br />\n         4. Integrated persistence mechanisms tailored to each operating system<br />\n         5. Conducted extensive testing in isolated environments to validate stealth and functionality<br />\n         6. Documented cross-platform behavior and detection evasion capabilities</p>\n      \n      <h2>Key Takeaways</h2>\n      <p>— The dual-nature approach (media playback + background execution) significantly increases deception effectiveness<br />\n         — Polymorphic techniques successfully bypass standard antivirus and EDR solutions<br />\n         — Cross-platform persistence requires OS-specific implementation strategies<br />\n         — User behavior analysis confirms high success rates with media-based social engineering<br />\n         — Thumbnail and metadata manipulation dramatically improve click-through rates</p>\n      \n      <h2>Technical Implementation</h2>\n      <p>The executable utilizes advanced process hollowing techniques to simultaneously handle media decoding and payload execution. The polymorphic engine regenerates code signatures upon each execution, while maintaining core functionality. Persistence is achieved through various methods including scheduled tasks, launch agents, and service installation depending on the target platform.</p>\n      \n      <h2>Community Impact & Future Work</h2>\n      <p>This research highlights critical gaps in how systems handle seemingly legitimate media executables. The project demonstrates that traditional media file analysis overlooks executable-based threats that incorporate media content. Future enhancements could include AI-driven polymorphism, cloud-based payload updates, and expanded social engineering tactics.</p>\n      \n      <h2>Personal Reflection</h2>\n      <p>Developing this threat vector revealed the delicate balance between functionality and malice. The technical challenge of maintaining seamless media playback while executing undetectable code required innovative approaches to process management and resource allocation. This knowledge directly contributes to improving defensive strategies against similarly sophisticated attacks.</p>\n    </div>", "summary": "A security research project detailing the development of a polymorphic malware disguised as media files, featuring cross-platform persistence, FUD capabilities, and social engineering tactics for a hackathon demonstration.", "published_on": "2025-09-25", "category": "Offensive Security", "tags": ["malware-development", "red-teaming", "polymorphic-code", "persistence", "social-engineering", "fud-malware", "cross-platform"], "author": "b33lz3bubTH", "read_time": "4 min read"}' > /dev/null

# Create Journal Entry 4
echo "Creating Journal Entry 4..."
curl -s -X POST "$BASE_URL/api/journal" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d '{"title": "ATS-Optimized Resume Generator", "body": "<div class='\''vintage-content'\''>\n  <h2>Overview</h2>\n  <p>An AI-driven pipeline that generates an ATS-friendly, professional resume tailored to a specific job description. Built on top of <a href=\"https://github.com/b33lz3bubTH/pixel-perfect-resume\" target=\"_blank\" rel=\"noopener noreferrer\">pixel-perfect-resume</a>, one script run produces a polished PDF ready for submission.</p>\n\n  <h2>How it works</h2>\n  <ol>\n    <li><strong>AI Analysis:</strong> A large language model analyzes your existing resume and the target job description to extract key skills, responsibilities, and target keywords.</li>\n    <li><strong>Optimization:</strong> The content is rewritten and reordered to maximize keyword relevance and ATS parsing accuracy while preserving clarity for human reviewers.</li>\n    <li><strong>Formatting:</strong> Optimized content is mapped into clean, ATS-compliant templates and professional themes.</li>\n    <li><strong>PDF Generation:</strong> The final result is exported as a print-ready PDF that maintains structure and readability across applicant tracking systems.</li>\n  </ol>\n\n  <h2>Why it matters</h2>\n  <p>Applicant Tracking Systems often filter candidates before a human ever sees the resume. This tool focuses on both machine-readability and human-friendly presentation so your application has the best chance to pass automated screening and impress recruiters.</p>\n\n  <h2>Get started</h2>\n  <p>Run the script in the linked repository to generate your ATS-optimized PDF in seconds. For more details and customization options, visit the <a href=\"https://github.com/b33lz3bubTH/pixel-perfect-resume\" target=\"_blank\" rel=\"noopener noreferrer\">project on GitHub</a>.</p>\n</div>", "summary": "AI-powered resume generator that analyzes a job description and your resume, optimizes content for ATS compatibility, applies professional formatting, and outputs a print-ready PDF.", "published_on": "2023-10-17"}' > /dev/null

# Create Meme Categories and capture IDs
echo "Creating Meme Categories..."
CATEGORY1=$(curl -s -X POST "$BASE_URL/api/memes/categories" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d '{"name": "souravOldLinuxRicing"}' | grep -o '"id":"[^"]*"' | cut -d'"' -f4)

CATEGORY2=$(curl -s -X POST "$BASE_URL/api/memes/categories" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d '{"name": "darkjokes"}' | grep -o '"id":"[^"]*"' | cut -d'"' -f4)

CATEGORY3=$(curl -s -X POST "$BASE_URL/api/memes/categories" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d '{"name": "LetMeBrag"}' | grep -o '"id":"[^"]*"' | cut -d'"' -f4)

# Add Memes to souravOldLinuxRicing
echo "Adding memes to souravOldLinuxRicing..."
curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY1\", \"type\": \"img\", \"src\": \"https://github.com/b33lz3bubTH/dotfiles/raw/main/images/qtile.png\"}" > /dev/null

curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY1\", \"type\": \"img\", \"src\": \"https://github.com/b33lz3bubTH/dotfiles/raw/main/images/nvim.png\"}" > /dev/null

curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY1\", \"type\": \"img\", \"src\": \"https://github.com/b33lz3bubTH/dotfiles/raw/main/images/rofi_1.png\"}" > /dev/null

curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY1\", \"type\": \"img\", \"src\": \"https://github.com/b33lz3bubTH/dotfiles/raw/main/images/rofi_2.png\"}" > /dev/null

curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY1\", \"type\": \"img\", \"src\": \"https://github.com/b33lz3bubTH/dotfiles/raw/main/images/xmonad/xmonad.png\"}" > /dev/null

curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY1\", \"type\": \"img\", \"src\": \"https://github.com/b33lz3bubTH/dotfiles/raw/main/images/xmonad/doom.png\"}" > /dev/null

curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY1\", \"type\": \"img\", \"src\": \"https://github.com/b33lz3bubTH/dotfiles/raw/main/images/xmonad/fileexp.png\"}" > /dev/null

curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY1\", \"type\": \"img\", \"src\": \"https://github.com/b33lz3bubTH/dotfiles/raw/main/images/xmonad/kitty.png\"}" > /dev/null

# Add Memes to darkjokes
echo "Adding memes to darkjokes..."
curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY2\", \"type\": \"img\", \"src\": \"https://i.chzbgr.com/full/9751564288/h39525B80\"}" > /dev/null

curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY2\", \"type\": \"img\", \"src\": \"https://i.chzbgr.com/full/9751564544/hC0033DEE\"}" > /dev/null

curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY2\", \"type\": \"webm\", \"src\": \"https://img-9gag-fun.9cache.com/photo/a1mqmew_460swp.webp\"}" > /dev/null

curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY2\", \"type\": \"webm\", \"src\": \"https://img-9gag-fun.9cache.com/photo/aByNZEP_460swp.webp\"}" > /dev/null

curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY2\", \"type\": \"webm\", \"src\": \"https://img-9gag-fun.9cache.com/photo/aLnNLeA_460swp.webp\"}" > /dev/null

# Add Memes to LetMeBrag
echo "Adding memes to LetMeBrag..."
curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY3\", \"type\": \"yt\", \"src\": \"https://www.youtube.com/watch?v=ALOqMiXNL70\"}" > /dev/null

curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY3\", \"type\": \"yt\", \"src\": \"https://www.youtube.com/watch?v=K-6RoM2pTB4\"}" > /dev/null

curl -s -X POST "$BASE_URL/api/memes" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d "{\"category_id\": \"$CATEGORY3\", \"type\": \"yt\", \"src\": \"https://www.youtube.com/watch?v=77VyA6NTgn8\"}" > /dev/null

# Create Stories
echo "Creating Stories..."
curl -s -X POST "$BASE_URL/api/stories" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d '{"media": "https://pbs.twimg.com/media/GTwE1JmbwAA3nOI?format=jpg&name=large", "mimetype": "image/jpeg", "title": "I Love Cats <3", "description": "Choccssssssss are my life"}' > /dev/null

curl -s -X POST "$BASE_URL/api/stories" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d '{"media": "https://pbs.twimg.com/media/Gt3udmSXkAAfrOR?format=jpg&name=large", "mimetype": "image/jpeg", "title": "I love long rides, i love heavy bikes, i love metal both music and bikes.", "description": "but sm650 doesn'\''t matter now."}' > /dev/null

curl -s -X POST "$BASE_URL/api/stories" \
  -H "Authorization: Bearer $ROOT_KEY" \
  -H "Content-Type: application/json" \
  -d '{"media": "https://pbs.twimg.com/media/F_N6FVbagAAajGS?format=jpg&name=4096x4096", "mimetype": "image/jpeg", "title": "Coffee Break", "description": "Taking a break with my coding companion"}' > /dev/null

echo "Seeding completed!"

