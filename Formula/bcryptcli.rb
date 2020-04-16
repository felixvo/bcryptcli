# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Bcryptcli < Formula
  desc "Bcrypt CLI"
  homepage "https://github.com/felixvo/homebrew-bcrypt"
  url "https://github.com/felixvo/homebrew-bcrypt/raw/master/bcrypt.zip"
  version "1.0.0"
  sha256 "ccedb8cd031512ca24c52a5c323b6b61c5c6c893a2a99d07604c9be4cfb222d9"

  # depends_on "cmake" => :build

  def install
    # ENV.deparallelize  # if your formula fails when building in parallel
    # Remove unrecognized options if warned by configure
    bin.install "bcrypt"
    # system "cmake", ".", *std_cmake_args
  end

  test do
    # `test do` will create, run in and delete a temporary directory.
    #
    # This test will fail and we won't accept that! For Homebrew/homebrew-core
    # this will need to be a test that verifies the functionality of the
    # software. Run the test with `brew test bcryptcli`. Options passed
    # to `brew install` such as `--HEAD` also need to be provided to `brew test`.
    #
    # The installed folder is not in the path, so use the entire path to any
    # executables being tested: `system "#{bin}/program", "do", "something"`.
    system "false"
  end
end
