'use client'

import { useState, useEffect, useRef } from 'react'
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Search, Menu } from 'lucide-react'

export function CoursesPageComponent() {
  const [isNavVisible, setIsNavVisible] = useState(false)
  const navRef = useRef(null)
  const coursesRef = useRef([])

  useEffect(() => {
    const handleMouseMove = (event) => {
      if (event.clientY < 50) {
        setIsNavVisible(true)
      } else {
        setIsNavVisible(false)
      }
    }

    document.addEventListener('mousemove', handleMouseMove)

    const observerOptions = {
      root: null,
      rootMargin: '0px',
      threshold: 0.1,
    }

    const observerCallback = (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.classList.add('expanded')
        } else {
          entry.target.classList.remove('expanded')
        }
      })
    }

    const observer = new IntersectionObserver(observerCallback, observerOptions)

    coursesRef.current.forEach((course) => {
      if (course) observer.observe(course)
    })

    return () => {
      document.removeEventListener('mousemove', handleMouseMove)
      coursesRef.current.forEach((course) => {
        if (course) observer.unobserve(course)
      })
    };
  }, [])

  const courses = [
    { id: 1, title: 'Introduction to React', description: 'Learn the basics of React and build your first app.' },
    { id: 2, title: 'Advanced JavaScript', description: 'Deep dive into advanced JavaScript concepts and patterns.' },
    { id: 3, title: 'CSS Mastery', description: 'Master CSS and create beautiful, responsive layouts.' },
    { id: 4, title: 'Node.js Fundamentals', description: 'Build server-side applications with Node.js and Express.' },
    { id: 5, title: 'Python for Data Science', description: 'Use Python for data analysis and machine learning.' },
    { id: 6, title: 'UI/UX Design Principles', description: 'Learn the fundamentals of creating user-friendly interfaces.' },
    { id: 7, title: 'Mobile App Development', description: 'Create cross-platform mobile apps using React Native.' },
    { id: 8, title: 'DevOps and CI/CD', description: 'Implement efficient development and deployment pipelines.' },
    { id: 9, title: 'Blockchain Basics', description: 'Understand the fundamentals of blockchain technology.' },
    { id: 10, title: 'Machine Learning Fundamentals', description: 'Get started with machine learning algorithms and applications.' },
    { id: 11, title: 'Cloud Computing with AWS', description: 'Learn to build and deploy applications on Amazon Web Services.' },
    { id: 12, title: 'Cybersecurity Essentials', description: 'Understand key concepts in protecting digital assets and information.' },
  ]

  return (
    (<div className="min-h-screen bg-gray-100">
      <nav
        ref={navRef}
        className={`fixed top-0 left-0 right-0 bg-white shadow-md transition-transform duration-300 ease-in-out z-10 ${
          isNavVisible ? 'translate-y-0' : '-translate-y-full'
        }`}>
        <div className="container mx-auto px-4 py-4 flex justify-between items-center">
          <h1 className="text-xl font-bold">Course Catalog</h1>
          <Button variant="ghost" size="icon">
            <Menu className="h-6 w-6" />
            <span className="sr-only">Menu</span>
          </Button>
        </div>
      </nav>
      <main className="container mx-auto px-4 pt-20">
        <h1 className="text-4xl font-bold text-center my-10">Explore Our Courses</h1>

        <div className="max-w-2xl mx-auto my-16">
          <div className="relative">
            <Search
              className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
            <Input
              type="search"
              placeholder="Search for courses..."
              className="w-full pl-10 pr-4 py-3 text-lg rounded-lg border border-gray-300 focus:border-blue-500 focus:ring-2 focus:ring-blue-200" />
          </div>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mt-16">
          {courses.map((course, index) => (
            <Card
              key={course.id}
              ref={(el) => (coursesRef.current[index] = el)}
              className="transform transition-all duration-300 ease-in-out hover:scale-105 opacity-70 hover:opacity-100 flex flex-col">
              <CardHeader>
                <CardTitle className="text-lg">{course.title}</CardTitle>
                <CardDescription className="text-sm">{course.description}</CardDescription>
              </CardHeader>
              <CardContent className="mt-auto">
                <Button className="w-full">Learn More</Button>
              </CardContent>
            </Card>
          ))}
        </div>
      </main>
    </div>)
  );
}