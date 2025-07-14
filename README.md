 Master The Go Programming Language: Elevate Your Skills!

 Go Bootcamp with gRPC and Protocol Buffers

Welcome to the ultimate journey into Go programming! Whether you're a beginner aiming to dive headfirst into software development or an experienced coder looking to sharpen your skills, this comprehensive course is tailored just for you.

Important Note
🧠 Why Do I Get "main redeclared" Errors in Go? Understanding package main, main() Functions, and Folder Structure in Your Go Course
If you're seeing an error like:

main redeclared in this block
Or you're wondering why you can’t go run a file after changing the package name, you’re not alone. Many students encounter this while following along with the course, so let’s break it down step-by-step and clarify how Go handles files, packages, and the main function.

✅ The Setup in This Course
In this course, we build understanding one file at a time — first hello_world.go, then data_types.go, variables.go, and so on.

But only one file exists in the main folder at a time.

Once a concept is covered in a lecture, that file is either:

🔄 Renamed (e.g., to .txt), so Go won’t try to compile it.

📁 Moved into a subfolder like basics/, and its package line is updated to match the folder (package basics).

This is done intentionally to keep the learning experience clean and prevent errors like multiple main() declarations.

🧨 Why Does This Error Happen?
Let’s say you have two files in the same folder:

random.go

package main
 
func main() {
    fmt.Println("Random code")
}
variables.go

package main
 
func main() {
    fmt.Println("Variables code")
}
You will get an error:

main redeclared in this block
This is because Go allows only ONE main() function per package, and by default, all .go files in a folder belong to the same package (unless explicitly specified otherwise).

💡 What is package main?
In Go, if you're writing an executable program, it must:

Use package main

Have a main() function as the entry point

If you try to run a file without package main or without a main() function, you’ll get an error like:

go run: cannot run non-main package
💡 Then What is package basics, package utils, etc.?
These are library-style packages — used for reusable code. They don’t need a main() function, and they’re not meant to be executed directly with go run.

That’s why, in some lectures, I change the file from:

package main
to:

package basics
Before archiving or moving it — this tells Go, "Hey, this is not a program to run anymore."

🛠️ Your Options
Here’s how you can avoid or fix these errors:

✅ Option 1: Follow the folder and naming structure exactly as I am doing in the lectures
✅ Option 2: Keep Only One main() at a Time
If you want to keep everything in one folder (e.g., during early learning), just make sure only one file has a main() function.

You can temporarily rename other .go files to .txt or comment out extra main() functions.

✅ Option 3: Use Subfolders for Different Programs
Structure your code like this:

/Go_course/
  /random/
    random.go   --> package main (has its own main())
  /variables/
    variables.go --> package main (has its own main())
Now you can cd into each folder and run:

go run random.go
Each file is part of a different Go program now. No conflicts!

✅ Option 4: Combine Logic Into One File
Instead of having two main() functions, split functionality into named functions:

package main
 
func main() {
    sayHello()
    showVariables()
}
 
func sayHello() {
    fmt.Println("Hello from Random")
}
 
func showVariables() {
    fmt.Println("Hello from Variables")
}
This is good for practice and helps organize your code better.

🤔 Why Do I Rename Files or Change the Package?
When you see me renaming files or changing package main to package basics, it’s to avoid compilation errors. This way, I can keep the previous lecture’s code as a reference without Go thinking I’m trying to run multiple programs at once.

⚠️ Important Rules to Remember
✅ Allowed                                                                             |                 ❌ Not Allowed

----------------------------------------------------------------------------------------------------------------------

One main() in a package Multiple                                   |                 main() functions in the same package

Files in different folders, each with package main       |                  Two main() files in the same folder

go run file.go with package main                              | go run on a file with a different package name

Archiving old .go files by changing package or extension |  Leaving many .go files with main() in one folder

🔁 FAQ Summary
Q: Can I have multiple .go files with main()?
A: Yes, but only if they’re in separate folders or packages.

Q: Why change the package name to basics?
A: To indicate that the file is now part of a library, not meant to be run.

Q: Why does go run give an error after changing the package name?
A: Because only package main can be run as an executable.

Q: Why does VS Code show squiggly lines when I have multiple files?
A: Because it detects multiple main() functions or duplicate declarations.



💡 Why I’m Only Changing the Package Name (Not the main() Function)

You might have noticed that inside the basics, intermediate, and advanced folders, many files still have a main() function — and that’s completely okay. I don’t rename the main() function in those files because we’re not trying to run them directly while they’re sitting in those folders. Each of those folders is like a “storage area” for code examples we’ve already covered. If you tried to run any of those folders directly, you’d get an error — because Go doesn’t allow multiple files with main() functions in the same package (i.e., folder) to be compiled together.

Instead of spending time removing or renaming every main() function, I simply change the package name to match the folder (package basics, package intermediate, etc.). That way, the files are still saved for reference, but Go won’t treat them as executable programs anymore — it will treat them like regular library code that’s not meant to be run directly.

Later, if you want to run any of those example files again, just move the file out into the main folder, rename its package to main, and make sure there’s no other file with a main() function in that folder. Then you can run it normally using:

go run your_file.go
This setup helps us keep things clean, focused, and organized by topic, without getting stuck on Go’s restrictions around main() functions.



As we progress into the course and as you grow more comfortable, you’ll learn how to structure projects more professionally (and there is an exclusive coverage on this as well, later on), but right now, the goal is understanding Go fundamentals — not full-scale architecture.

Stay curious, keep experimenting, and don’t worry — you’re doing great! 🚀

Happy coding!
— Ashish