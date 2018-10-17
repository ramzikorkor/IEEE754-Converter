# IEEE754-Converter

The IEEE Standard for Floating-Point Arithmetic (IEEE 754) is a technical standard for floating-point computation established in 1985 by the Institute of Electrical and Electronics Engineers (IEEE). The standard addressed many problems found in the diverse floating point implementations that made them difficult to use reliably and portably. Many hardware floating point units now use the IEEE 754 standard.

The standard defines:

arithmetic formats: sets of binary and decimal floating-point data, which consist of finite numbers (including signed zeros and subnormal numbers), infinities, and special "not a number" values (NaNs)
interchange formats: encodings (bit strings) that may be used to exchange floating-point data in an efficient and compact form
rounding rules: properties to be satisfied when rounding numbers during arithmetic and conversions
operations: arithmetic and other operations (such as trigonometric functions) on arithmetic formats
exception handling: indications of exceptional conditions (such as division by zero, overflow, etc.)
The current version, IEEE 754-2008 published in August 2008, includes nearly all of the original IEEE 754-1985 standard and the IEEE Standard for Radix-Independent Floating-Point Arithmetic (IEEE 854-1987).

^^^^^Source Wikipedia^^^^^^

Senior year in my degree at Boise State University, we were tasked with using IEEE 754 for class. I knew that I could just look up one of the numerous converters online, but I challenged myself to create my own. I figured if I ever needed it again I could make a libarary that would make it very simple. Here is my take on converting IEEE 754 single and double precision. 


