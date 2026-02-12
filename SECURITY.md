# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 6.x.x   | :white_check_mark: |
| 5.x.x   | :white_check_mark: |
| < 5.0   | :x:                |

## Reporting a Vulnerability

If you discover a security vulnerability within MathsWithGolang, please follow these steps:

1. **Do NOT** create a public GitHub issue for security vulnerabilities
2. Email the maintainer directly at: **mouaadidoufkir.contact@gmail.com**
3. Include a detailed description of the vulnerability
4. Provide steps to reproduce the issue if possible

### What to expect

- **Acknowledgment**: We will acknowledge receipt of your report within 48 hours
- **Investigation**: We will investigate and validate the reported vulnerability
- **Resolution**: We aim to resolve critical vulnerabilities within 7 days
- **Disclosure**: We will coordinate with you on public disclosure timing

## Security Best Practices

When using this library:

1. **Input Validation**: Always validate inputs before passing to mathematical functions
2. **Overflow Protection**: Be aware of integer/float overflow in extreme calculations
3. **Precision Limits**: Understand floating-point precision limitations
4. **Resource Limits**: Set appropriate limits for iterative algorithms to prevent infinite loops

## Code Security

This library:

- ✅ Uses **100% pure Golang** with no external dependencies
- ✅ Has **no network calls** or external data fetching
- ✅ Contains **no file system operations** beyond what Go stdlib provides
- ✅ Is designed for **deterministic mathematical computations**

## Acknowledgments

We appreciate responsible disclosure of security issues to help keep MathsWithGolang safe for everyone.

---

**Created by: MOUAAD IDOUFKIR**
